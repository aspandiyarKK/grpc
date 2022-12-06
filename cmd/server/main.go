package main

import (
	"google.golang.org/grpc"
	"grpc/proto"
	"grpc/server"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	proto.RegisterXrServer(s, srv)
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
