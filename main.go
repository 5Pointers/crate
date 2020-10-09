package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/5Pointers/crate/storage"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9999, "The server port")
)

type storageServer struct {
	storage.UnimplementedDataStorageServiceServer
}

func (s *storageServer) Set(ctx context.Context, entry *storage.DataEntry) (*storage.DataEntry, error) {
	log.Println("Incoming data!")
	log.Println(entry)
	return entry, nil
}

// func (s *storageServer) Find(query *storage.Query, stream *storage.DataStorageService_FindServer) error {
// 	return nil
// }

func getServer() *storageServer {
	s := &storageServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	storage.RegisterDataStorageServiceServer(grpcServer, getServer())
	grpcServer.Serve(lis)
}
