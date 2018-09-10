package engine

import (
	"crontab_worker/config"
	"crontab_worker/scheduler"
	"strconv"
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

func createWorker(in chan config.Job, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			job := <-in
			err := DoWork(job)
			if err != nil {
				continue
			}
		}
	}()
}
