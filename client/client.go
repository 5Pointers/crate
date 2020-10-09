package main

import (
	"context"
	"log"
	"time"

	"github.com/5Pointers/crate/storage"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := storage.NewDataStorageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Set(ctx, &storage.DataEntry{Payload: map[string]string{"name": "john"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Response:")
	log.Println(r)
	log.Println("===")
}
