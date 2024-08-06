package pkg

import (
	"fmt"
	"medal-service/config"

	"medal-service/internal/service"
	"net"

	pb "github.com/Bekzodbekk/protofiles/genproto/medals"

	"google.golang.org/grpc"
)

type Service struct {
	medalService *service.MedalService
}

func NewService(medalService *service.MedalService) *Service {
	return &Service{
		medalService: medalService,
	}
}

func (s *Service) RUN(cfg config.Config) error {
	target := fmt.Sprintf("%s:%s", cfg.MedalServiceHost, cfg.MedalServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterMedalServiceServer(serv, s.medalService)

	return serv.Serve(listener)
}
