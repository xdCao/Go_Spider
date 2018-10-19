package main

import (
	"Go_Spider/crawler/engine"
	"Go_Spider/crawler/scheduler"
	"Go_Spider/crawler/zhenai/parser"
)

func main() {

	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedSchduler{},
		WorkerCount: 100,
	}
	//
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc:parser.ParseCity,
	//})

}
