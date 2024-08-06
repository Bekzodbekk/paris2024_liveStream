package userDial

import (
	"api-service/config"
	"fmt"

	"github.com/Bekzodbekk/protofiles/genproto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	user.UserServiceClient
}

func DialWithUserService(conf config.Config) (*UserService, error) {
	target := fmt.Sprintf("%s:%s", conf.AuthServiceHost, conf.AuthServicePort)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userService := user.NewUserServiceClient(client)

	return &UserService{UserServiceClient: userService}, nil
}
