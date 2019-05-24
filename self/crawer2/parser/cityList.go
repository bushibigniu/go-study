package parser

import (
	"fmt"
	"go-study/self/crawer2/engine"
	"regexp"
)
//<a target="_blank" href="http://www.zhenai.com/zhenghun/shenzhen">深圳征婚</a>
const regCity  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


func ParseCityList(content []byte)  engine.ParseResult{

	//1.regexp
	re := regexp.MustCompile(regCity)
	matchs := re.FindAllSubmatch(content, -1)

	//2. 将获取的数据放到一个集合里面
	result := engine.ParseResult{}
	for _,m := range matchs{

		//item 地名
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]), //城市链接
			ParserFunc:engine.NilParse,
		})

		fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
	}

	fmt.Print("matchs len :", len(matchs))

	return result

}
