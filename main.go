package main

import (
	"log"
	"net"

	"github.com/jihadable/learn-grpc-category/service"
	categoryPb "github.com/jihadable/learn-grpc-proto/category"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	categoryService := service.NewCategoryService()
	categoryPb.RegisterCategoryServiceServer(grpcServer, categoryService)

	log.Println("Category gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
