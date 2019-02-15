package main

import (
	"imooc.com/u2pppw/crawler/engine"
	"imooc.com/u2pppw/crawler/persist"
	"imooc.com/u2pppw/crawler/scheduler"
	"imooc.com/u2pppw/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
