package scheduler

import "go-crawler/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedule) WorkerReady(chan engine.Request) {

}

func (s *SimpleSchedule) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedule) Submit(r engine.Request) {
	//s.workerChan <- r
	go func() {s.workerChan <- r}()
}
