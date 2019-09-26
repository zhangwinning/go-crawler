package main

import (
	"go-crawler/engine"
	"go-crawler/zhenai/parser"
)

func main() {
	result := engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	}
	//fmt.Printf("%s", result.ParseFunc)
	engine.Run(result)
}
