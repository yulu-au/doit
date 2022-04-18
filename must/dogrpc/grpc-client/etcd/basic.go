package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var EClient *clientv3.Client

func init() {
	var err error
	EClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
}

func example() {
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	_, err := EClient.Put(ctx, "sample_key", "sample_value")

	if err != nil {
		panic(err)
	}
}
