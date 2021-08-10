package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"orderArchive/api"
)

func main() {

	server := grpc.NewServer()
	instance := &api.OrderServer{}
	api.RegisterOrderServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8881")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}

}
