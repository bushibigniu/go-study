package main

import (
	"go-study/self/crawer2/engine"
	"go-study/self/crawer2/parser"
)

func main ()  {

	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	})
}
