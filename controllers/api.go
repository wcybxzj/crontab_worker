package controllers

import (
	"crontab_worker/config"
	"crontab_worker/engine"
	"crontab_worker/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ReceiveConfigedJob(c *gin.Context) {
	var res models.ResponseData
	var jobReq config.Job
	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			res.Code = strconv.Itoa(models.ERROR_JOB_WORLING)
			res.Msg = "fail, this JobId: " + jobReq.JobId + " is running!"
			c.JSON(http.StatusUnauthorized, res)
			return
		}

		//任务是否是配置好的
		for _, jobData := range config.Config.Jobs {
			fmt.Println(jobData.JobId)
			fmt.Println(jobReq.JobId)
			if jobReq.JobId == jobData.JobId {
				engine.E.Scheduler.Submit(jobData)
				res.Code = strconv.Itoa(models.OK)
				res.Msg = "ok"
				c.JSON(http.StatusOK, res)
				return
			}
		}

		res.Code = strconv.Itoa(models.ERROR_UNKNOWN_JOBID)
		res.Msg = "fail, the JobId not exsits in worker config file"
		c.JSON(http.StatusUnauthorized, res)
	}
}

func ReceiveDiyJob(c *gin.Context) {
	var res models.ResponseData
	var jobReq config.Job
	if c.BindJSON(&jobReq) == nil {
		ok := engine.IsJobWorking(jobReq.JobId)
		if ok == true {
			res.Code = strconv.Itoa(models.ERROR_JOB_WORLING)
			res.Msg = "fail, this JobId: " + jobReq.JobId + " is running!"
			c.JSON(http.StatusUnauthorized, res)
			return
		}

		engine.E.Scheduler.Submit(jobReq)
		res.Code = strconv.Itoa(models.OK)
		res.Msg = "ok"
		c.JSON(http.StatusOK, res)

	}
}

//Reload api
func Reload(c *gin.Context) {
	var res models.ResponseData

	if c.ClientIP() != "127.0.0.1" {
		res.Code = strconv.Itoa(models.ERROR_REQ_FORMAT)
		res.Msg = "Request error!"
		c.JSON(http.StatusMethodNotAllowed, res)
		return
	}

	config.LoadConfig()

	res.Code = strconv.Itoa(models.OK)
	res.Msg = "ok!"
	c.JSON(http.StatusOK, res)
}
