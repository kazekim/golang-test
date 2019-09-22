/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"context"
	"fmt"
	"github.com/kazekim/golang-test/testviper/remotestore/model"
	"github.com/kazekim/golang-test/testviper/remotestore/security"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/xdefrag/viper-etcd/remote"
	etcd "go.etcd.io/etcd/client"
)

func init() {
	viper.RemoteConfig = &remote.Config{
		Decoder: &model.Decode{},
	}
}

// You can start ETCD container like this:
// docker run -d -p 2379:2379 quay.io/coreos/etcd:latest etcd --advertise-client-urls http://0.0.0.0:2380 --listen-client-urls http://0.0.0.0:2379

func main() {
	if os.Getenv("ETCD_ADDR") == "" {
		must(os.Setenv("ETCD_ADDR", "http://0.0.0.0:2379"))
	}

	initEtcdKeys()

	vpr := viper.New()

	must(vpr.AddRemoteProvider("etcd", os.Getenv("ETCD_ADDR"), "/jirawatkim"))

	vpr.SetConfigType("json")

	must(vpr.ReadRemoteConfig())

	var wg sync.WaitGroup
	wg.Add(1)

	var d model.Decode

	go func() {
		for {
			time.Sleep(time.Minute)
			must(vpr.WatchRemoteConfig())
			must(vpr.Unmarshal(&d))
			fmt.Println(d)
		}
	}()

	must(vpr.Unmarshal(&d))
	fmt.Println(d)

	wg.Wait()

}

func initEtcdKeys() {
	client, err := etcd.New(etcd.Config{
		Endpoints: []string{os.Getenv("ETCD_ADDR")},
	})
	if err != nil {
		panic(err)
	}

	kapi := etcd.NewKeysAPI(client)

	ctx := context.Background()

	must2(kapi.Set(ctx, "/jirawatkim/access/token", security.DoEncode("some_token"), nil))
	must2(kapi.Set(ctx, "/jirawatkim/redis/addr", security.DoEncode("http://0.0.0.0:6379"), nil))
	must2(kapi.Set(ctx, "/jirawatkim/redis/password", security.DoEncode("veryStrongPassword"), nil))
	must2(kapi.Set(ctx, "/jirawatkim/deeply/nested/config/wow", security.DoEncode("this_is_value"), nil))
	must2(kapi.Set(ctx, "/jirawatkim/providers", security.DoEncode([]string{"redis", "postgres"}), nil))
	must2(kapi.Set(ctx, "/jirawatkim/lucky/numbers", security.DoEncode([]int{9, 13}), nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func must2(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}