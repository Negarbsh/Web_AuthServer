package main

import (
	"AuthServer/service"
	"context"
)

type server struct {
	UnimplementedAuthServer
}

func (s *server) ReqPq(ctx context.Context, input *ReqPqInput) (*ReqPqResponse, error) {
	pg := service.GetPg(input.Nonce, input.MessageId)
	return &ReqPqResponse{
		Nonce:       pg.Nonce,
		ServerNonce: pg.ServerNonce,
		MessageId:   int32(pg.MessageId),
		P:           int32(pg.P),
		G:           int32(pg.G),
	}, nil
}

func (s *server) Req_DHParams(ctx context.Context, input *Req_DHParamsInput) (*Req_DHParamsResponse, error) {
	dh := service.GetDHParams(input.Nonce, input.ServerNonce, int(input.MessageId), int(input.A))
	return &Req_DHParamsResponse{
		Nonce:       dh.Nonce,
		ServerNonce: dh.ServerNonce,
		MessageId:   int32(dh.MessageId),
		B:           int32(dh.PublicKey),
	}, nil
}
