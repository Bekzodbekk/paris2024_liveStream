package handlers

import (
	"github.com/Bekzodbekk/protofiles/genproto/medals"
	"github.com/Bekzodbekk/protofiles/genproto/user"
)

type Handlers struct {
	UserService  user.UserServiceClient
	MedalService medals.MedalServiceClient
}
