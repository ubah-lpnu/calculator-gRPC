package server

import (
	"context"
	"log"

	calc "github.com/ubah-lpnu/calculator-gRPC/internal/calculate"
	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
)

type CalculateServer struct {
	pb.UnimplementedCalculateServer
}

func (s *CalculateServer) PerformCalculation(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Printf("Performing calculation for A=%f, B=%f", req.A, req.B)

	result, err := calc.MultiplyAndDivide(req.A, req.B)
	if err != nil {
		log.Printf("Calculation failed: %v", err)
		return nil, err
	}

	log.Printf("Calculation successful: ResultMul=%f, ResultDiv=%f", result.ResultMul, result.ResultDiv)

	return &pb.CalculateResponse{
		ResultMul: result.ResultMul,
		ResultDiv: result.ResultDiv,
	}, nil
}

func (s *CalculateServer) mustEmbedUnimplementedCalculateServer() {
	panic("unimplemented")
}
