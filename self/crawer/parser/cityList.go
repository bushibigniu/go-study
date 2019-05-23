package parser

import (
	"go-study/self/crawer/engine"
	"regexp"
)

const cityListRe  = `<a href="http://www.zhenai.com/[0-9a-z]+"[^>]*>[^<]+</a>`


func ParseCityList(content []byte)  engine.ParseResult {
	//rex := `<a href="http://www.zhenai.com/[0-9a-z]+"[^>]*>[^<]+</a>`


	re := regexp.MustCompile(cityListRe)
	//matchs := re.FindAll(content, -1)
	matchs := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _,m := range matchs{
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:string(m[1]),
				//ParseFunc:nil,
				ParseFunc:engine.NilParser,
			})
	}

	return result
}