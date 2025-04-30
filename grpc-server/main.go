package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "grpc-server/helloworld"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":50051"
	httpPort = ":8080"
)

// server is used to implement helloworld.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// startHealthCheckServer starts an HTTP server for health checks.
func startHealthCheckServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("HTTP health check server listening on %s", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatalf("failed to start HTTP health check server: %v", err)
	}
}

func main() {
	// Start health check server in a goroutine
	go startHealthCheckServer()

	// Start gRPC server
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
} 