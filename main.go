package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleSchedule{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
