package scheduler

import "Go_Spider/crawler/engine"

type SimpleScheduler struct {
	workerCHan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerCHan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerCHan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerCHan <- request
	}()
}
