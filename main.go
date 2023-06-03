package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("gRPC auth server")

	listener, err := net.Listen("tcp", ":5052")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	RegisterAuthServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
