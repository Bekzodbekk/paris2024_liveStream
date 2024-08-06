package service

import (
	"context"
	"user-service/internal/repository"

	pb "github.com/Bekzodbekk/protofiles/genproto/user"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepository
	rdb      *redis.Client
}

func NewService(user repository.UserRepository, rdb *redis.Client) *UserService {
	return &UserService{
		userRepo: user,
		rdb:      rdb,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.userRepo.Register(ctx, req)
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.userRepo.Login(ctx, req)
}

func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return s.userRepo.RefreshToken(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.userRepo.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.userRepo.DeleteUser(ctx, req)
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.userRepo.GetUserById(ctx, req)
}

func (s *UserService) GetUsers(ctx context.Context, req *pb.Void) (*pb.GetUsersResponse, error) {
	return s.userRepo.GetUsers(ctx, req)
}

func (s *UserService) GetUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.GetUsersResponse, error) {
	return s.userRepo.GetUserByFilter(ctx, req)
}
