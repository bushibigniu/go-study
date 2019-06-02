package main

import (
	"go-study/self/16dataAndView/simpleitemServer/engine"
	"go-study/self/16dataAndView/simpleitemServer/parser"
	"go-study/self/16dataAndView/simpleitemServer/scheduler"
)
/**
	scheduler 实现第一步：所有worker 公用一个输入

 */
func main()  {

	url := "http://www.zhenai.com/zhenghun"

	//simple
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:url,
	//	ParserFunc:parser.ParseCityList,
	//})

	//它是指针接收者，需要定个变量

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount:10, //没定这个会没反应
	}
	e.Run(engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	})
}
