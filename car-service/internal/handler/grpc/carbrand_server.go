package grpc

import (
	"context"
	"ecommerce/car-service/internal/domain"
	"ecommerce/car-service/internal/pb/proto"
	"ecommerce/car-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CarBrandServer struct {
	pb.UnimplementedCarBrandServiceServer
	uc usecase.CarBrandUsecase
}

func NewCarBrandServer(uc usecase.CarBrandUsecase) *CarBrandServer {
	return &CarBrandServer{uc: uc}
}

// ListCarBrands RPC
func (s *CarBrandServer) ListCarBrands(ctx context.Context, req *pb.ListCarBrandsRequest) (*pb.ListCarBrandsResponse, error) {
	carbrands, err := s.uc.List()
	if err != nil {
		return nil, err
	}
	pbCarBrands := []*pb.CarBrand{}
	for _, cb := range carbrands {
		pbCarBrands = append(pbCarBrands, &pb.CarBrand{
			Id:          cb.ID,
			Name:        cb.Name,
			Description: cb.Description,
		})
	}
	return &pb.ListCarBrandsResponse{Carbrands: pbCarBrands}, nil
}

// GetCarBrand RPC
func (s *CarBrandServer) GetCarBrand(ctx context.Context, req *pb.GetCarBrandRequest) (*pb.CarBrand, error) {
	cb, err := s.uc.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	if cb == nil {
		return nil, status.Errorf(codes.NotFound, "car brand not found")
	}
	return &pb.CarBrand{
		Id:          cb.ID,
		Name:        cb.Name,
		Description: cb.Description,
	}, nil
}

// CreateCarBrand RPC
func (s *CarBrandServer) CreateCarBrand(ctx context.Context, req *pb.CreateCarBrandRequest) (*pb.CarBrand, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}
	brand := &domain.CarBrand{
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.uc.Create(brand)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create car brand: %v", err)
	}
	return &pb.CarBrand{
		Id:          brand.ID, // Only valid if your usecase or handler sets this!
		Name:        brand.Name,
		Description: brand.Description,
	}, nil

}

// UpdateCarBrand RPC
func (s *CarBrandServer) UpdateCarBrand(ctx context.Context, req *pb.UpdateCarBrandRequest) (*pb.CarBrand, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}
	brand := &domain.CarBrand{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.uc.Update(brand)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update car brand: %v", err)
	}
	return &pb.CarBrand{
		Id:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
	}, nil
}

// DeleteCarBrand RPC
func (s *CarBrandServer) DeleteCarBrand(ctx context.Context, req *pb.DeleteCarBrandRequest) (*pb.DeleteCarBrandResponse, error) {
	err := s.uc.Delete(req.Id)
	success := err == nil
	return &pb.DeleteCarBrandResponse{Success: success}, err
}
