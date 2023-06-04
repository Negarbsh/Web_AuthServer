package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:5052")
	if err != nil {
		panic(err)
	}
	fmt.Println("gRPC auth server running on " + listener.Addr().String())

	s := grpc.NewServer()
	RegisterAuthServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
