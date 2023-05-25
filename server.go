package main

import (
	"context"
	"fmt"
)

type server struct {
	UnimplementedAuthServer
}

func (s *server) ReqPq(ctx context.Context, input *ReqPqInput) (*ReqPqResponse, error) {
	fmt.Println(input.Nonce)
	return &ReqPqResponse{
		Nonce:       "sslam",
		ServerNonce: "sdfsdfsd",
		MessageId:   12323,
		P:           334,
		G:           23,
	}, nil
}

func (s *server) Req_DHParams(ctx context.Context, input *Req_DHParamsInput) (*Req_DHParamsResponse, error) {
	fmt.Println(input.MessageId)
	return &Req_DHParamsResponse{
		Nonce:       "sdfsdf",
		ServerNonce: "sfsdf",
		MessageId:   123123,
		B:           33,
	}, nil
}
