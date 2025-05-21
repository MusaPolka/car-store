package grpc

import (
	"context"
	"ecommerce/order-service/internal/pb/proto"
	"ecommerce/order-service/internal/usecase"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	uc usecase.OrderUsecase
}

func NewOrderServer(uc usecase.OrderUsecase) *OrderServer {
	return &OrderServer{uc: uc}
}

// Example: ListOrders implementation
func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.uc.List()
	if err != nil {
		return nil, err
	}
	pbOrders := []*pb.Order{}
	for _, o := range orders {
		// Convert domain.Order to pb.Order
		pbOrder := &pb.Order{
			Id:         o.ID,
			UserId:     o.UserID,
			TotalPrice: o.TotalPrice,
			Status:     o.Status,
		}
		// You may want to populate pbOrder.Items here if you have items in the domain model.
		pbOrders = append(pbOrders, pbOrder)
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
