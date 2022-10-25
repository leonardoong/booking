package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoong/booking/internal/common/resource"
	"github.com/leonardoong/booking/internal/config"
)

func newRoutes(cfg *config.Config, res *resource.Resources, usecase Usecase) *gin.Engine {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	return router
}
