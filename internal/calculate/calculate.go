package calculate

import (
	"errors"

	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
)

func MultiplyAndDivide(a, b float32) (*pb.CalculateResponse, error) {
	result := &pb.CalculateResponse{}
	resultMulChan := make(chan float32)
	resultDivChan := make(chan float32)

	if b == 0 {
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
