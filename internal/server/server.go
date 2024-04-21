package server

import (
	"context"

	calc "github.com/ubah-lpnu/calculator-gRPC/internal/calculate"
	"github.com/ubah-lpnu/calculator-gRPC/internal/logger"
	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
	"go.uber.org/zap"
)

type CalculateServer struct {
	pb.UnimplementedCalculateServer
}

func (s *CalculateServer) PerformCalculation(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log := logger.GetLogger()
	log.Info("performing calculation",
		zap.Float64("a", req.A),
		zap.Float64("b", req.B),
	)

	result, err := calc.MultiplyAndDivide(req.A, req.B)
	if err != nil {
		log.Error("calculation failed", zap.Error(err))
		return nil, err
	}

	log.Info("calculation successful",
		zap.Float64("ResultMul", result.ResultMul),
		zap.Float64("ResultDiv", result.ResultDiv),
	)

	return &pb.CalculateResponse{
		ResultMul: result.ResultMul,
		ResultDiv: result.ResultDiv,
	}, nil
}

func (s *CalculateServer) mustEmbedUnimplementedCalculateServer() {
	panic("unimplemented")
}
