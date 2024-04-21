package main

import (
	"log"
	"net"

	"github.com/ubah-lpnu/calculator-gRPC/internal/server"
	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	s := grpc.NewServer()
	srv := &server.CalculateServer{}
	pb.RegisterCalculateServer(s, srv)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
