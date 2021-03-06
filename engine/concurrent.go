package engine

import (
	"crontab_worker/config"
	"crontab_worker/scheduler"
	"strconv"
	"zuji/common/dlog"
)

var E ConcurrentEngine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(job config.Job)
	WorkerChan() chan config.Job
	Run()
	QueueStatus() string
}

type ReadyNotifier interface {
	WorkerReady(chan config.Job)
}

func init_concurrent() {
	maxGoroutines, err := strconv.Atoi(config.Config.MaxGoroutines)
	if err != nil {
		panic(err)
	}
	E = ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: maxGoroutines,
	}
}

func (e *ConcurrentEngine) Run() {
	init_concurrent()
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), e.Scheduler)
	}
}

func (e *ConcurrentEngine) QueueStatus() string {
	return e.QueueStatus()
}

func createWorker(in chan config.Job, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			job := <-in
			dlog.LogColor(dlog.TextGreen, "DoWork start, JobId:"+job.JobId)
			err := DoWork(job)
			if err != nil {
				dlog.LogColor(dlog.TextRed, "DoWork err:"+err.Error()+", JobId:"+job.JobId)
				continue
			}
			dlog.LogColor(dlog.TextGreen, "DoWork finish JobId:"+job.JobId)
		}
	}()
}
