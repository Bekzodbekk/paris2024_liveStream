package main

import (
	"log"
	"user-service/config"
	"user-service/internal/service"
	"user-service/internal/storage"
	"user-service/pkg"
	"user-service/postgres"
	"user-service/redis"
)

func main() {
	conf := config.Load()

	db, err := postgres.InitDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	rds, err := redis.ConnectRedis(conf)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := storage.NewUserRepo(db, rds)
	service := service.NewService(userRepo, rds)
	runServ := pkg.NewService(*service)

	log.Printf("User service running on :%s port", conf.UserServicePort)
	if err := runServ.RUN(conf); err != nil{
		log.Fatal(err)
	}
}
