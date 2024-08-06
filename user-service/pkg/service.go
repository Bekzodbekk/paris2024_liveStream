package pkg

import (
	"fmt"
	"net"
	"user-service/config"
	"user-service/internal/service"

	pb "github.com/Bekzodbekk/protofiles/genproto/user"

	"google.golang.org/grpc"
)

type Service struct {
	userService service.UserService
}

func NewService(userService service.UserService) *Service {
	return &Service{
		userService: userService,
	}
}

func (s *Service) RUN(cfg config.Config) error {
	target := fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterUserServiceServer(serv, &s.userService)

	return serv.Serve(listener)
}
