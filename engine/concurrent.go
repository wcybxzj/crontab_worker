package engine

import (
	"crontab_worker/config"
	"crontab_worker/global"
	"crontab_worker/scheduler"
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

func init() {
	E = ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: global.MAX_GOROUTINES,
	}
}
func (e *ConcurrentEngine) Run() {
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
