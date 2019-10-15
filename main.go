package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	//result := engine.Request{
	//	URL:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//}
	////fmt.Printf("%s", result.ParseFunc)
	//engine.SimpleEngine{}.Run(result)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleSchedule{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
