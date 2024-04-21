package calculate

import (
	"errors"

	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
)

func MultiplyAndDivide(a, b float32) (*pb.CalculateResponse, error) {
	result := &pb.CalculateResponse{}
	resultChan := make(chan float32)

	if b == 0 {
		return nil, errors.New("Division by 0")
	}

	go func() {
		resultChan <- a * b
	}()

	go func() {
		resultChan <- a / b
	}()

	result.ResultMul = <-resultChan
	result.ResultDiv = <-resultChan

	return result, nil

}
