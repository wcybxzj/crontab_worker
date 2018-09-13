package main

import (
	"crontab_worker/config"
	"crontab_worker/controllers"
	"crontab_worker/engine"
	"github.com/gin-gonic/gin"
	"zuji/common/dlog"
)

func main() {
	config.LoadConfig()

	dlog.OpenLogfile("data/crontab_worker.log")
	defer dlog.CloseLogfile()

	engine.E.Run()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/ReceiveConfigedJob", controllers.ReceiveConfigedJob)
	router.POST("/ReceiveDiyJob", controllers.ReceiveDiyJob)
	router.Any("/Check", controllers.Check)
	router.Any("/Reload", controllers.Reload)
	router.Any("/QueueStatus", controllers.QueueStatus)
	router.Run(":8080")
}
