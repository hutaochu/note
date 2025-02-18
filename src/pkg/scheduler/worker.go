package scheduler

import "fmt"

type Job struct {
	Name     string
	priority int
}

type Worker struct {
	queue chan *Job
	stop  chan struct{}
	state string
}

func (w *Worker) Start() {
	for {
		select {
		case <-w.stop:
			return
		case j := <-w.queue:
			fmt.Println(j.Name + " is done")
		}
	}
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}

type Scheduler struct {
	workerPool []Worker
	workQueue  chan *Job
	jobPool    []Job
}

func NewScheduler(workerCount int) *Scheduler {
	return &Scheduler{
		workerPool: make([]Worker, workerCount),
		workQueue:  make(chan *Job, workerCount),
	}
}

func (s *Scheduler) Add(j Job) {
	s.workQueue <- &j
}

func (s *Scheduler) Run() {
	go func() {
		//for j := range s.workQueue {
		//
		//}
	}()
}

type WorkerPool struct{}
