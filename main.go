package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// TODO: load configs from env
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("unable to listen server: %v", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc server not started: %s", err)
	}
}
