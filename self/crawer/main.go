package main

import (
	"go-study/self/crawer/engine"
	"go-study/self/crawer/parser"
)

func main()  {

	url := "http://www.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url: url,
		ParseFunc:parser.ParseCityList,
	})
}