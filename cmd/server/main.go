package main

import (
	"fmt"
	"net"

	log "github.com/ubah-lpnu/calculator-gRPC/internal/logger"
	"github.com/ubah-lpnu/calculator-gRPC/internal/server"
	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	Port = ":50051"
)

func main() {
	log := log.GetLogger()
	log.Info("Server started on port", zap.String("port", Port))
	s := grpc.NewServer()
	srv := &server.CalculateServer{}
	pb.RegisterCalculateServer(s, srv)

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}

}
