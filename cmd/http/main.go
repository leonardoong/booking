package main

import (
	"fmt"
	"log"

	"github.com/leonardoong/booking/internal/config"
	"github.com/leonardoong/booking/internal/constant"
)

const (
	repoName = constant.ServiceBookingSystem
	service  = constant.ServiceHTTP
	appName  = repoName + "-" + service
)

func main() {
	fmt.Println("Booking System BE HTTP Server")

	// init config
	cfg, err := config.New(appName, repoName)
	if err != nil {
		log.Fatalf("failed to init config: %v \n", err)
	}

	// start http app
	err = startApp(cfg)
	if err != nil {
		log.Fatalf("failed to start app: %v \n", err)
	}

}
