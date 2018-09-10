package main

import (
	"crontab_worker/config"
	"crontab_worker/controllers"
	"crontab_worker/engine"
	"github.com/gin-gonic/gin"
)

func main() {
	engine.E.Run()
	config.LoadConfig()
	router := gin.Default()
	router.POST("/ReceiveConfigedJob", controllers.ReceiveConfigedJob)
	router.POST("/ReceiveDiyJob", controllers.ReceiveDiyJob)
	router.Any("/check", controllers.Check)
	router.Any("/reload", controllers.Reload)
	router.Run(":8080")
}
