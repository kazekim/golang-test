/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/pkg/errors"
	"context"
	pb "github.com/kazekim/golang-test/grpctls/proto"
)

var (
	host    = "localhost"
	port    = "50051"
	address = host + ":" + port
)

const (
	tlsNoVerify = iota + 1
	tlsVerifyNoCA
	tlsVerifyWithCA
	tlsWithCertFile
	insecureNoTLS
	defaultConfig
)

type input struct {
	skipVerify bool
	cert       string
	ca         string
}

type tlsCredsFunc func(i input) ([]grpc.DialOption, error)

func tlsCredsAuto(i input) ([]grpc.DialOption, error) {
	config := &tls.Config{
		InsecureSkipVerify: i.skipVerify,
	}
	return []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(config))}, nil
}

func tlsCredsAutoCA(i input) ([]grpc.DialOption, error) {
	b, err := ioutil.ReadFile(i.ca)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, errors.New("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: i.skipVerify,
		RootCAs:            cp,
	}
	return []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(config))}, nil
}

func tlsCredsFile(i input) ([]grpc.DialOption, error) {
	creds, err := credentials.NewClientTLSFromFile(i.cert, "")
	if err != nil {
		return nil, errors.Wrap(err, "could not process the credentials")
	}
	return []grpc.DialOption{grpc.WithTransportCredentials(creds)}, nil
}

func noTLS(i input) ([]grpc.DialOption, error) {
	return []grpc.DialOption{grpc.WithInsecure()}, nil
}

func defaultTLS(i input) ([]grpc.DialOption, error) {
	config := &tls.Config{}
	return []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(config))}, nil
}

func main() {
	id := flag.Uint("id", 1, "User ID")
	cafile := flag.String("file", "ca.cert", "CA public certificate")
	mode := flag.Uint("mode", 1, "User ID")
	flag.Parse()

	var f tlsCredsFunc
	var i input

	switch *mode {
	case tlsNoVerify:
		i = input{skipVerify: true}
		f = tlsCredsAuto
	case tlsVerifyNoCA:
		i = input{skipVerify: false}
		f = tlsCredsAuto
	case tlsVerifyWithCA:
		i = input{ca: *cafile, skipVerify: false}
		f = tlsCredsAutoCA
	case tlsWithCertFile:
		i = input{cert: "service.pem"}
		f = tlsCredsFile
	case insecureNoTLS:
		i = input{}
		f = noTLS
	case defaultConfig:
		i = input{}
		f = defaultTLS
		address = host + ":443"
	}

	ctx := context.Background()
	opts, err := f(i)
	if err != nil {
		log.Fatalf("problem creating the tls credentials: %v", err)
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGUserClient(conn)
	res, err := client.GetByID(ctx, &pb.GetByIDRequest{Id: uint32(*id)})
	if err != nil {
		log.Fatalf("Server says: %v", err)
	}
	fmt.Println("User found: ", res.GetName())
}