package main

import (
	"log"
	"medal-service/config"
	"medal-service/internal/service"
	"medal-service/internal/storage"
	"medal-service/pkg"
	"medal-service/postgres"
)

func main() {
	conf := config.Load()

	db, err := postgres.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	medalRepo := storage.NewMedalRepo(db)
	service := service.NewMedalService(medalRepo)

	serviceRun := pkg.NewService(service)
	log.Printf("Medal service running on :%s port", conf.MedalServicePort)
	if err := serviceRun.RUN(conf); err != nil {
		log.Fatal(err)
	}
}
