package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	grpcHandler "ecommerce/order-service/internal/handler/grpc"
	"ecommerce/order-service/internal/infrastructure/postgres"
	pb "ecommerce/order-service/internal/pb/proto"
	repo "ecommerce/order-service/internal/repository/postgres"
	"ecommerce/order-service/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewDBConn("localhost", "postgres", "P@ssw0rd", "ecom", 5432)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	orderRepo := repo.NewOrderRepo(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, grpcHandler.NewOrderServer(orderUsecase))
	reflection.Register(grpcServer)

	log.Println("gRPC order service running at :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
