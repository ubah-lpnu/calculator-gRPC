package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ubah-lpnu/calculator-gRPC/pkg/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	a := flag.Float64("a", 0.0, "first decimal number")
	b := flag.Float64("b", 0.0, "second decimal number")

	flag.Parse()

	if *a == 0.0 || *b == 0.0 {
		log.Fatal("provide both numbers")
	}

	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()
	client := pb.NewCalculateClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.PerformCalculation(ctx, &pb.CalculateRequest{A: *a, B: *b})
	if err != nil {
		log.Fatalf("could not perform calculation: %v", err)
	}

	log.Printf("Result of multiplication: %f", res.ResultMul)
	log.Printf("Result of division: %f", res.ResultDiv)

}
