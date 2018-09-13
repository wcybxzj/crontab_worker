package controllers

import (
	"crontab_worker/config"
	"crontab_worker/engine"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"strconv"
	"zuji/common/dlog"
)

func ReceiveConfigedJob(c *gin.Context) {
	var jobReq config.Job

	err := c.BindJSON(&jobReq)
	if err != nil {
		dlog.LogColor(dlog.TextRed, "fail BindJSON", err)
		return
	}

	ok := engine.IsJobWorking(jobReq.JobId)
	if ok == true {
		dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
		config.Output(c)
		return
	}

	for _, jobData := range config.Config.Jobs {
		if jobReq.JobId == jobData.JobId {
			engine.E.Scheduler.Submit(jobData)
			dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
			config.Output(c)
			return
		}
	}

	dlog.LogColor(dlog.TextRed, "fail, the JobId not exsits in worker config file")
	config.Output(c)
	return

}

func ReceiveDiyJob(c *gin.Context) {
	var jobReq config.Job

	err := c.BindJSON(&jobReq)
	if err != nil {
		dlog.LogColor(dlog.TextRed, "fail BindJSON, err:", err)
		return
	}

	ok := engine.IsJobWorking(jobReq.JobId)
	if ok == true {
		dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
		config.Output(c)
		return
	}

	engine.E.Scheduler.Submit(jobReq)
	dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
	config.Output(c)
	return

}

//Reload api
func Reload(c *gin.Context) {
	if c.ClientIP() != "127.0.0.1" {
		dlog.LogColor(dlog.TextGreen, "fail, only allow 127.0.0.1")
		config.Output(c)
		return
	}

	config.LoadConfig()
	config.Output(c)
	return
}

func QueueStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"goroutine num": strconv.Itoa(runtime.NumGoroutine()),
			"QueueStatus":   engine.E.Scheduler.QueueStatus(),
		},
	})
}
