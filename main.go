package main

import (
	"concurrent_crawler/engine"
	"concurrent_crawler/scheduler"
	"concurrent_crawler/zhenai/parser"
)


func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 5,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
