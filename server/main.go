package main

import (
	"context"
	"log"
	"net"

	pb "github.com/thinhnguyen-com/calculator/calculatorrpc"
	"google.golang.org/grpc"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.GetA() + req.GetB()
	return &pb.AddResponse{Result: result}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := req.GetA() - req.GetB()
	return &pb.SubtractResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &calculatorServer{})

	log.Println("🚀 Server listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
