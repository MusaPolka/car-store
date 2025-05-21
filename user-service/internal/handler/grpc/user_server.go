package grpc

import (
	"context"
	"ecommerce/user-service/internal/pb/proto"
	"ecommerce/user-service/internal/usecase"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	uc usecase.UserUsecase
}

func NewUserServer(uc usecase.UserUsecase) *UserServer {
	return &UserServer{uc: uc}
}

// Example: ListUsers implementation
func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.uc.List()
	if err != nil {
		return nil, err
	}
	pbUsers := []*pb.User{}
	for _, u := range users {
		pbUser := &pb.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			FullName: u.FullName,
		}
		pbUsers = append(pbUsers, pbUser)
	}
	return &pb.ListUsersResponse{Users: pbUsers}, nil
}
