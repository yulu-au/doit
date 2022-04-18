package main

import (
	"context"
	"dogrpc/pb"
	"dogrpc/service"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var port = flag.String("port", "8080", "please type port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("we listen on %v\n", lis.Addr())
	var opts []grpc.ServerOption
	opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp), grpc.ChainUnaryInterceptor(UnaryInterceptor1))
	grpcServer := grpc.NewServer(opts...)
	// grpcServer := grpc.NewServer(
	// 	// grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	// 	// 	grpc_ctxtags.StreamServerInterceptor(),
	// 	// 	grpc_opentracing.StreamServerInterceptor(),
	// 	// 	grpc_prometheus.StreamServerInterceptor,
	// 	// 	grpc_zap.StreamServerInterceptor(zapLogger),
	// 	// 	grpc_auth.StreamServerInterceptor(myAuthFunction),
	// 	// 	grpc_recovery.StreamServerInterceptor(),
	// 	// )),
	// 	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	// 		grpc_ctxtags.UnaryServerInterceptor(),
	// 		grpc_opentracing.UnaryServerInterceptor(),
	// 		grpc_prometheus.UnaryServerInterceptor,
	// 		// grpc_zap.UnaryServerInterceptor(zapLogger),
	// 		// grpc_auth.UnaryServerInterceptor(myAuthFunction),
	// 		grpc_recovery.UnaryServerInterceptor(),
	// 	)),
	// )
	reflection.Register(grpcServer)
	pb.RegisterGreeterServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() pb.GreeterServer {
	return &service.Greeter{}
}

func UnaryInterceptor1(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	var ip string
	p, ok := peer.FromContext(ctx)
	if ok {
		ip = p.Addr.String()
	}
	md, _ := metadata.FromIncomingContext(ctx)
	start := time.Now()
	resp, err = handler(ctx, req)
	end := time.Now()
	log.Printf("in1 %10s | %14s | %10v | md=%v | reply = %v", ip, info.FullMethod, end.Sub(start), md, resp)
	return
}
