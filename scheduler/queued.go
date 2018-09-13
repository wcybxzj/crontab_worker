package scheduler

import (
	"crontab_worker/config"
	"fmt"
)

type QueuedScheduler struct {
	jobChan    chan config.Job
	workerChan chan chan config.Job
	jobQ       []config.Job
	workerQ    []chan config.Job
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
		for {
			var activeJob config.Job
			var activeWorker chan config.Job

			if len(s.jobQ) > 0 && len(s.workerQ) > 0 {
				activeJob = s.jobQ[0]
				activeWorker = s.workerQ[0]
			}

			select {
			case j := <-s.jobChan:
				s.jobQ = append(s.jobQ, j)

			case w := <-s.workerChan:
				s.workerQ = append(s.workerQ, w)

			case activeWorker <- activeJob:
				s.jobQ = s.jobQ[1:]
				s.workerQ = s.workerQ[1:]
			}
		}
	}()
}

func (s *QueuedScheduler) QueueStatus() string {
	return fmt.Sprintf("workerQ中空闲worker:%d, JobQ中积压待处理job:%d", len(s.workerQ), len(s.jobQ))
}
