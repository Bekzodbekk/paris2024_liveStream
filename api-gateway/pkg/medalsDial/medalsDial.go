package medalsdial

import (
	"api-service/config"
	"fmt"

	"github.com/Bekzodbekk/protofiles/genproto/medals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MedalService struct {
	medals.MedalServiceClient
}

func DialWithMedalService(conf config.Config) (*MedalService, error) {
	target := fmt.Sprintf("%s:%s", conf.MedalServiceHost, conf.MedalServicePort)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	medalService := medals.NewMedalServiceClient(client)

	return &MedalService{MedalServiceClient: medalService}, nil
}
