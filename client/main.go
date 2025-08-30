package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thinhnguyen-com/calculator/calculatorrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Dùng NewClient thay vì Dial
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test Add
	addResp, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add Result: %d", addResp.GetResult())

	// Test Subtract
	subResp, err := client.Subtract(ctx, &pb.SubtractRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtract Result: %d", subResp.GetResult())
}
