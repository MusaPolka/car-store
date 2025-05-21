package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	grpcHandler "ecommerce/user-service/internal/handler/grpc"
	"ecommerce/user-service/internal/infrastructure/postgres"
	pb "ecommerce/user-service/internal/pb/proto"
	repo "ecommerce/user-service/internal/repository/postgres"
	"ecommerce/user-service/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewDBConn("localhost", "postgres", "P@ssw0rd", "ecom", 5432)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	userRepo := repo.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcHandler.NewUserServer(userUsecase))
	reflection.Register(grpcServer)

	log.Println("gRPC user service running at :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
