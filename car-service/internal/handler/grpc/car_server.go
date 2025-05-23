package grpc

import (
	"context"
	"ecommerce/car-service/internal/domain"
	"ecommerce/car-service/internal/pb/proto"
	"ecommerce/car-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CarServer struct {
	pb.UnimplementedCarServiceServer
	uc usecase.CarUsecase
}

func NewCarServer(uc usecase.CarUsecase) *CarServer {
	return &CarServer{uc: uc}
}

// ListCars RPC
func (s *CarServer) ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.ListCarsResponse, error) {
	cars, err := s.uc.List(0, 0) // you can add paging if needed
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list cars: %v", err)
	}
	pbCars := []*pb.Car{}
	for _, car := range cars {
		pbCars = append(pbCars, toProtoCar(&car))
	}
	return &pb.ListCarsResponse{Cars: pbCars}, nil
}

// GetCar RPC
func (s *CarServer) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.Car, error) {
	car, err := s.uc.GetByID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get car: %v", err)
	}
	if car == nil {
		return nil, status.Errorf(codes.NotFound, "car not found")
	}
	return toProtoCar(car), nil
}

// CreateCar RPC
func (s *CarServer) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.Car, error) {
	newCar := &domain.Car{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		BrandID:     req.GetBrandId(),
		Price:       req.GetPrice(),
		Stock:       int(req.GetStock()),
		Year:        int(req.GetYear()),
		Color:       req.GetColor(),
	}
	err := s.uc.Create(newCar)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create car: %v", err)
	}
	return toProtoCar(newCar), nil
}

// DeleteCar RPC
func (s *CarServer) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteCarResponse, error) {
	err := s.uc.Delete(req.GetId())
	if err != nil {
		return &pb.DeleteCarResponse{Success: false}, status.Errorf(codes.Internal, "failed to delete car: %v", err)
	}
	return &pb.DeleteCarResponse{Success: true}, nil
}

// Helper to map domain.Car -> pb.Car
func toProtoCar(car *domain.Car) *pb.Car {
	return &pb.Car{
		Id:          car.ID,
		Name:        car.Name,
		Description: car.Description,
		BrandId:     car.BrandID,
		Price:       car.Price,
		Stock:       int32(car.Stock),
		Year:        int32(car.Year),
		Color:       car.Color,
	}
}
