package repository

import (
	"context"

	pb "github.com/Bekzodbekk/protofiles/genproto/user"
)

type UserRepository interface {
	Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	GetUsers(ctx context.Context, req *pb.Void) (*pb.GetUsersResponse, error)
	GetUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.GetUsersResponse, error)
}
