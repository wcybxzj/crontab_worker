package controllers

import (
	"crontab_worker/config"
	"crontab_worker/engine"
	"github.com/gin-gonic/gin"
	"net/http"
	"zuji/common/dlog"
)

//{"code":"0","msg":"\u6210\u529f","status":"ok","data":[]}
func output(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		//"data": gin.H{
		//	"addJob": gin.H{"name": "crontab_worker"},
		//	"runJob": gin.H{"name": "crontab_worker"},
		//}
	})
}

func ReceiveConfigedJob(c *gin.Context) {
	var jobReq config.Job
	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
			output(c)
			return
		}

		//任务是否是配置好的
		for _, jobData := range config.Config.Jobs {
			if jobReq.JobId == jobData.JobId {
				engine.E.Scheduler.Submit(jobData)
				dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
				output(c)
				return
			}
		}

		dlog.LogColor(dlog.TextRed, "fail, the JobId not exsits in worker config file")
		output(c)
		return
	}
}

func ReceiveDiyJob(c *gin.Context) {
	var jobReq config.Job

	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			dlog.LogColor(dlog.TextRed, "fail, the JobId:"+jobReq.JobId+" is working")
			output(c)
			return
		}

		engine.E.Scheduler.Submit(jobReq)
		dlog.LogColor(dlog.TextGreen, "ok, the JobId:"+jobReq.JobId)
		output(c)
		return
	}
}

//Reload api
func Reload(c *gin.Context) {
	if c.ClientIP() != "127.0.0.1" {
		dlog.LogColor(dlog.TextGreen, "fail, only allow 127.0.0.1")
		output(c)
		return
	}

	config.LoadConfig()
	output(c)
	return
}
