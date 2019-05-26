package parser

import (
	"fmt"
	"go-study/self/15bingfa/crawer4/engine"
	"regexp"
)

//<a href="http://www.zhenai.com/zhenghun/aomen" data-v-5e16505f="">澳门</a>
const reCity = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//城市列表解析器
func ParseCityList(content []byte) engine.ParseResult {

	re := regexp.MustCompile(reCity)

	//matchs 这样的一个个集合，matchs[0],matchs[1],matchs[2]
	// [<a href="http://www.zhenai.com/zhenghun/zigong" data-v-5e16505f>
	// 自贡</a> http://www.zhenai.com/zhenghun/zigong 自贡]
	matchs := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	limitCity := 300
	for _, m := range matchs {
		//这边强转 为string, 防止类型不对
		//item 地名
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]), //城市链接
			ParserFunc: ParseUser,
		})

		limitCity--
		if limitCity == 0 {
			break
		}

	}

	fmt.Print("matchs len :", len(matchs))
	return result
}
