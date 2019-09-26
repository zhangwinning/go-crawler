package main

import (
	"go-crawler/engine"
	"go-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
