package main

import (
	"go-study/self/15bingfa/crawer2/engine"
	"go-study/self/15bingfa/crawer2/parser"
)
/**
	scheduler 实现第一步：所有worker 公用一个输入

 */
func main()  {

	url := "http://www.zhenai.com/zhenghun"

	//simple
	engine.SimpleEngine{}.Run(engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	})
}
