package controllers

import (
	"crontab_worker/config"
	"crontab_worker/engine"
	"github.com/gin-gonic/gin"
	"net/http"
	"zuji/common/dlog"
)

func ReceiveConfigedJob(c *gin.Context) {
	var jobReq config.Job
	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		//任务是否是配置好的
		for _, jobData := range config.Config.Jobs {
			if jobReq.JobId == jobData.JobId {
				engine.E.Scheduler.Submit(jobData)
				dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
				return
			}
		}

		dlog.LogColor(dlog.TextRed, "fail, the JobId not exsits in worker config file")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
}

func ReceiveDiyJob(c *gin.Context) {
	var jobReq config.Job

	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		engine.E.Scheduler.Submit(jobReq)
		dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
}

//Reload api
func Reload(c *gin.Context) {
	if c.ClientIP() != "127.0.0.1" {
		c.JSON(http.StatusOK, gin.H{"status": "fail, only allow 127.0.0.1"})
		return
	}

	config.LoadConfig()
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}
