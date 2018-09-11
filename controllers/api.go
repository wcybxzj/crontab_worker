package controllers

import (
	"crontab_worker/config"
	"crontab_worker/engine"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func buildRes(status string) config.ResponseData {
	var res config.ResponseData
	res.Status = status
	res.Data = config.Data{
		config.AddJob{config.JOB_NAME},
		config.RunJob{config.JOB_NAME},
	}
	return res
}

func ReceiveConfigedJob(c *gin.Context) {
	var jobReq config.Job
	var res config.ResponseData
	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			res = buildRes("fail, the JobId:" + jobReq.JobId + " is working")
			c.JSON(http.StatusOK, res)
			return
		}

		//任务是否是配置好的
		for _, jobData := range config.Config.Jobs {
			fmt.Println(jobData.JobId)
			fmt.Println(jobReq.JobId)
			if jobReq.JobId == jobData.JobId {
				engine.E.Scheduler.Submit(jobData)
				res = buildRes("ok")
				c.JSON(http.StatusOK, res)
				return
			}
		}

		res = buildRes("fail, the JobId not exsits in worker config file")
		c.JSON(http.StatusOK, res)
	}
}

func ReceiveDiyJob(c *gin.Context) {
	var res config.ResponseData
	var jobReq config.Job

	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			res = buildRes("fail, the JobId:" + jobReq.JobId + " is working")
			c.JSON(http.StatusOK, res)
			return
		}

		engine.E.Scheduler.Submit(jobReq)
		res = buildRes("ok")
		c.JSON(http.StatusOK, res)
	}
}

//Reload api
func Reload(c *gin.Context) {
	var res config.ResponseData

	if c.ClientIP() != "127.0.0.1" {
		res = buildRes("fail, only allow 127.0.0.1")
		c.JSON(http.StatusOK, res)
		return
	}

	config.LoadConfig()
	res = buildRes("ok")
	c.JSON(http.StatusOK, res)
}
