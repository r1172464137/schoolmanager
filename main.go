package main

import (
	"github.com/gin-gonic/gin"
	"school_manager/configs"
	"school_manager/routes"
)

func main() {
	configs.Init()
	gin.SetMode(configs.AppMode)
	r := routes.NewRouter()
	r.Run(configs.HttpPort)
}
