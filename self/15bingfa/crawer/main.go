package main

import (
	"go-study/self/15bingfa/crawer/engine"
	"go-study/self/15bingfa/crawer/parser"
)

func main()  {

	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	})
}
