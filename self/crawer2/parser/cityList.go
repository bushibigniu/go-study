package parser

import (
	"go-study/self/crawer2/engine"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"regexp"
)

const regCity  = ""

func ParseCityList(content []byte)  engine.ParseResult{

	//1.regexp
	re := regexp.MustCompile(regCity)

	matchs := re.FindAllSubmatch(content, -1)

	//2. 将获取的数据放到一个集合里面

	result := engine.ParseResult{}
	for _,m := range matchs{
		result.Items = append(result.Items, m[2])

		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc:engine.NilParse,
		})

		fmt.Printf("City: %s, URL: %s", m[2], m[1])
	}

	fmt.Print("matchs len :", len(matchs))

	return result

}
