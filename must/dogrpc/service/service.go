package service

import (
	"context"
	"dogrpc/pb"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := pb.HelloReply{Message: in.Name}
	return &reply, nil
}

func (g *Greeter) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := pb.HelloReply{Message: in.Name}
	return &reply, nil
}
