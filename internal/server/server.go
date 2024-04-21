package server

import (
	"context"

	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
)

type CalculateServer struct {
	pb.UnimplementedCalculateServer
}

func (s *CalculateServer) PerformCalculation(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	mulResultChan := make(chan float32)
	divisionResult := make(chan float32)

	go func() {
		multiplication := req.A * req.B
		divisionResult <- multiplication
	}()

	go func() {
		division := req.A / req.B
		divisionResult <- division
	}()

	product := <-mulResultChan
	quotient := <-divisionResult

	return &pb.CalculateResponse{ResultMul: product, ResultDiv: quotient}, nil
}

func (s *CalculateServer) mustEmbedUnimplementedCalculateServer() {
	panic("unimplemented")
}
