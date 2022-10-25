package main

import (
	"log"

	"github.com/leonardoong/booking/internal/common/resource"
	"github.com/leonardoong/booking/internal/config"
)

func startApp(cfg *config.Config) error {
	// init resources
	res, err := resource.GetResources(cfg)
	if err != nil {
		log.Fatalf("failed to init resource: %v \n", err.Error())
	}

	usecase := initUsecase(cfg, res)
	router := newRoutes(cfg, res, usecase)

	router.Run(":8080")
	return nil
}
