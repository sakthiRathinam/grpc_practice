package main

import (
	"log"
	"net"

	pb "github.com/sakthiRathinam/explorerpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:9003"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v\n", err)
	}
}
