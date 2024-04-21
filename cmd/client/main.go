package main

import (
	"flag"
	"fmt"
	// "strings"
	// calc "github.com/ubah-lpnu/calculator-gRPC/internal/calculate"
	// pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
)

// func validateInputs([]inputNumbers string) (float32, float32, error) {
// 	numbers := strings.Split(inputNumbers, "")
// 	if len(numbers) != 2 {
// 		return 0, 0, errors.New("invalid input")
// 	}

// 	a, err := strconv.ParseFloat(numbers[0], 32)
// 	if err != nil {
// 		return 0, 0, err
// 	}
// 	b, err := strconv.ParseFloat(numbers[1], 32)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return float32(a), float32(b), nil
// }

func main() {

	a := flag.String("a", "1.0", "A number")
	b := flag.String("b", "1.0", "B number")
	flag.Parse()
	fmt.Println(*a, *b)
	fmt.Println(flag.Args())
}
