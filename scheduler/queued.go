package scheduler

import (
	"crontab_worker/config"
)

type QueuedScheduler struct {
	jobChan    chan config.Job
	workerChan chan chan config.Job
}

func (s *QueuedScheduler) WorkerChan() chan config.Job {
	return make(chan config.Job)
}

func (s *QueuedScheduler) Submit(j config.Job) {
	s.jobChan <- j
}

func (s *QueuedScheduler) WorkerReady(w chan config.Job) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan config.Job)
	s.jobChan = make(chan config.Job)

	go func() {
		var jobQ []config.Job
		var workerQ []chan config.Job

		for {
			var activeJob config.Job
			var activeWorker chan config.Job

			if len(jobQ) > 0 && len(workerQ) > 0 {
				activeJob = jobQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case j := <-s.jobChan:
				jobQ = append(jobQ, j)

			case w := <-s.workerChan:
				workerQ = append(workerQ, w)

			case activeWorker <- activeJob:
				jobQ = jobQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
