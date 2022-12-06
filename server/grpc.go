package server

import (
	"context"
	"grpc/proto"
	"math/rand"
)

type GRPCServer struct {
}

func (s *GRPCServer) GetCurrentRate(ctx context.Context, req *proto.CurrRequest) (*proto.Response, error) {
	return &proto.Response{Res: rand.Float32()}, nil
}

func (s *GRPCServer) Convert(ctx context.Context, req *proto.ConvertRequest) (*proto.Response, error) {
	return &proto.Response{Res: rand.Float32()}, nil
}
func (s *GRPCServer) GetHistoricalRate(ctx context.Context, req *proto.HistoricalRequest) (*proto.Response, error) {
	return &proto.Response{Res: rand.Float32()}, nil
}
func (s *GRPCServer) MustEmbedUnimplementedXrServer() {
	//TODO implement me
	panic("implement me")
}
