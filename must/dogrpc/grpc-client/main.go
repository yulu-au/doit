package main

import (
	"context"
	"fmt"
	"grpc-client/pb"
	"grpc-client/resolve"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	// _, err := etcd.EClient.Put(ctx, "sample_key", "sample_value")

	// if err != nil {
	// 	panic(err)
	// }
	resolver.Register(resolve.NewBuilder())
	conn, err := grpc.Dial("etcd:///echo", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cli := pb.NewEchoClient(conn)
	for i := 0; i < 1000; i++ {
		reply, err := cli.Echo(context.Background(), &pb.Echoreq{Msg: "freedom"})
		if err != nil {
			panic(err)
		}
		fmt.Println(reply.Msg)
		time.Sleep(time.Second)
	}
}
