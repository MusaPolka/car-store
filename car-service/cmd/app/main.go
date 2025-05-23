package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	grpcHandler "ecommerce/car-service/internal/handler/grpc" // <--- alias for your handler package
	db "ecommerce/car-service/internal/infrastructure/postgres"
	redisinfra "ecommerce/car-service/internal/infrastructure/redis"
	pb "ecommerce/car-service/internal/pb/proto"
	repo "ecommerce/car-service/internal/repository/postgres"
	"ecommerce/car-service/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.NewDBConn("localhost", "postgres", "P@ssw0rd", "ecom", 5432)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	cache := redisinfra.NewRedisClient("localhost:6379", "", 0)

	carBrandRepo := repo.NewCarBrandRepo(db)
	carBrandUsecase := usecase.NewCarBrandUsecase(carBrandRepo, cache) // <-- inject cache

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCarBrandServiceServer(grpcServer, grpcHandler.NewCarBrandServer(carBrandUsecase))
	reflection.Register(grpcServer)

	carRepo := repo.NewCarRepo(db)
	carUsecase := usecase.NewCarUsecase(carRepo)
	pb.RegisterCarServiceServer(grpcServer, grpcHandler.NewCarServer(carUsecase))

	log.Println("gRPC server running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
