package main

import (
	"api-service/api"
	"api-service/config"
	medalsdial "api-service/pkg/medalsDial"
	"api-service/pkg/userDial"
	"fmt"
	"log"
)

func main() {
	conf := config.Load()

	userDial, err := userDial.DialWithUserService(conf)
	if err != nil {
		log.Fatal(err)
	}

	medalDial, err := medalsdial.DialWithMedalService(conf)
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewGin(
		userDial,
		medalDial,
	)

	target := fmt.Sprintf("%s:%s", conf.ApiGatewayHost, conf.ApiGatewayPort)
	log.Printf("Api-gateway running on :%s port", conf.ApiGatewayPort)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
