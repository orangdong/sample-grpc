package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "grpc-client/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddress = getEnv("GRPC_ADDRESS", "https://mdeq7qigmp.ap-southeast-1.awsapprunner.com")
)

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return fallback
}

func main() {
	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set up a connection to the gRPC server
		conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("did not connect: %v", err)
			http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// Create a new client
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Make the gRPC call
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "HTTP Client"})
		if err != nil {
			log.Printf("could not greet: %v", err)
			http.Error(w, "Failed to get response from gRPC server", http.StatusInternalServerError)
			return
		}

		// Write the response back to the HTTP client
		fmt.Fprintf(w, "gRPC Response: %s", resp.Message)
	})

	log.Printf("HTTP server listening on :8080, connecting to gRPC server at %s", grpcAddress)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
} 