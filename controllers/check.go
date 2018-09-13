package controllers

import (
	"crontab_worker/config"
	"github.com/gin-gonic/gin"
)

//Check api
func Check(c *gin.Context) {
	config.Output(c)
}
