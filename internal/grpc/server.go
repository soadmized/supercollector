package grpc

import (
	"context"

	"supercollector/internal/pb"
)

type UserService interface {
}

type Server struct {
	pb.UnimplementedSupercollectorServer
}

// TODO
func (s Server) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	res := &pb.UserResponse{
		User: &pb.User{
			Name:    "John Snow",
			Email:   "john@snow.me",
			Pass:    "123pass",
			Phone:   "123455678",
			Comment: "comment",
		},
	}

	return res, nil
}
