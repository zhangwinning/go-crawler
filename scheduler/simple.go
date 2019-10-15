package scheduler

import "go-crawler/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) ConfigureMasterWorkerChan(c chan engine.Request) {
	//	send request down to worker chan
	s.workerChan = c
}

func (s *SimpleSchedule) Submit(r engine.Request) {
	//s.workerChan <- r
	go func() {s.workerChan <- r}()
}
