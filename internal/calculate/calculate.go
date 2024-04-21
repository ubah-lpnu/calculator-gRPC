package calculate

import (
	"errors"

	"github.com/ubah-lpnu/calculator-gRPC/internal/logger"
	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
	"go.uber.org/zap"
)

func MultiplyAndDivide(a, b float64) (*pb.CalculateResponse, error) {
	log := logger.GetLogger()
	result := &pb.CalculateResponse{}
	resultMulChan := make(chan float64)
	resultDivChan := make(chan float64)

	if b == 0 {
		log.Error("Division by zero", zap.Float64("A", a), zap.Float64("B", b))
		return nil, errors.New("Division by 0")
	}

	go func() {
		resultMulChan <- a * b
	}()

	go func() {
		resultDivChan <- a / b
	}()

	result.ResultMul = <-resultMulChan
	result.ResultDiv = <-resultDivChan

	return result, nil

}
