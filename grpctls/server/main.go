/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"context"
	"strings"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	pb "github.com/kazekim/golang-test/grpctls/proto"
	"errors"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
)

var (
	host   = "localhost"
	port   = "50051"
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
)

type userData struct {
	users map[uint32]string
}

func NewUserData() *userData {
	d := new(userData)
	d.users = make(map[uint32]string)
	d.users[1] = "Jirawat"
	return d
}

func (s *userData) GetByID(ctx context.Context, in *pb.GetByIDRequest) (*pb.User, error) {
	if s.users == nil {
		s.users = make(map[uint32]string)
	}
	if name, ok := s.users[in.Id]; ok {
		return &pb.User{
			Name: name,
			Id:   in.Id,
		}, nil
	}
	return nil, errors.New("user not found")
}

func testCert() (credentials.TransportCredentials, net.Listener) {
	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("golang-autocert"),
		HostPolicy: autocert.HostWhitelist(host),
		// Email: "",
	}

	return credentials.NewTLS(manager.TLSConfig()), manager.Listener()
}

type RSA struct {
	bits int
}

func (r RSA) Generate() (crypto.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, r.bits)
}

func vaultCert(f string) (credentials.TransportCredentials, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("vaultCert: problem with input file")
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("vaultCert: failed to append certificates")
	}
	issuer := &vault.Issuer{
		URL: &url.URL{
			Scheme: "https",
			Host:   "localhost:8200",
		},
		TLSConfig: &tls.Config{
			RootCAs: cp,
		},
		Token: "TOKEN HERE",
		Role:  "my-role",
	}
	cfg := certify.CertConfig{
		SubjectAlternativeNames: []string{"localhost"},
		IPSubjectAlternativeNames: []net.IP{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("::1"),
		},
		KeyGenerator: RSA{bits: 2048},
	}

	c := &certify.Certify{
		CommonName:  "localhost",
		Issuer:      issuer,
		Cache:       certify.NewMemCache(),
		CertConfig:  &cfg,
		RenewBefore: 24 * time.Hour,
	}
	tlsConfig := &tls.Config{
		GetCertificate: c.GetCertificate,
	}
	return credentials.NewTLS(tlsConfig), nil
}

func grpcHandlerFunc(g *grpc.Server, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if r.ProtoMajor == 2 && strings.Contains(ct, "application/grpc") {
			g.ServeHTTP(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func httpsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user from IP: %s\n\nYour config is: %+v", r.RemoteAddr, r.TLS)
	})
}

func main() {
	self := flag.Bool("self", true, "Whether to encrypt the connection using self-signed certs")
	public := flag.Bool("public", false, "Use certs emited by a trusted public CA")
	cefy := flag.Bool("cefy", false, "Use Certify with Vault as your CA")
	flag.Parse()

	logger = log.With(logger, "time", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	opts := []grpc.ServerOption{}
	var lis net.Listener
	var creds credentials.TransportCredentials
	var err error
	switch {
	// Public domain
	case *public:
		creds, lis = testCert()
		opts = append(opts, grpc.Creds(creds))
		// lis = autocert.NewListener(host)
		// port = "443"
	// Private domain
	default:
		switch {
		case *self && *cefy:
			_ = level.Error(logger).Log("msg", "can't choose self-signed and Certify at the same time")
			os.Exit(1)
		// Self-signed cetificates
		case *self:
			creds, err := credentials.NewServerTLSFromFile("service.pem", "service.key")
			if err != nil {
				_ = level.Error(logger).Log("msg", "failed to setup TLS with local files", "error", err)
				os.Exit(1)
			}
			opts = append(opts, grpc.Creds(creds))
		// Certificates signed by Vault via Certify
		case *cefy:
			creds, err := vaultCert("ca-org.cert")
			if err != nil {
				_ = level.Error(logger).Log("msg", "failed to setup TLS with Certify", "error", err)
				os.Exit(1)
			}
			opts = append(opts, grpc.Creds(creds))
		// Insecure
		default:
		}
		lis, err = net.Listen("tcp", ":"+port)
		if err != nil {
			_ = level.Error(logger).Log("msg", "failed to listen", "error", err)
			os.Exit(1)
		}
	}
	defer lis.Close()
	_ = level.Info(logger).Log("msg", "Server listening", "port", port)

	s := grpc.NewServer(opts...)
	_ = level.Info(logger).Log("msg", "Starting gRPC services")
	pb.RegisterGUserServer(s, NewUserData())

	_ = level.Info(logger).Log("msg", "Listening for incoming connections")
	if *public {
		if err = http.Serve(lis, grpcHandlerFunc(s, httpsHandler())); err != nil {
			_ = level.Error(logger).Log("msg", "failed to serve", "error", err)
			os.Exit(1)
		}
	} else {
		if err = s.Serve(lis); err != nil {
			_ = level.Error(logger).Log("msg", "failed to serve", "error", err)
			os.Exit(1)
		}
	}
}