package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}
	curr := flag.Arg(0)
	conn, err := grpc.Dial(":4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewXrClient(conn)
	res, err := client.GetCurrentRate(context.Background(), &proto.CurrRequest{Curr: curr})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetRes())
}
