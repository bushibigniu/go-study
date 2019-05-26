package parser

import (
	"go-study/self/15bingfa/crawer4/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1997037559" target="_blank">安静</a>
const regexpUser  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`


func ParseUser(body []byte)  engine.ParseResult{

	re := regexp.MustCompile(regexpUser)

	//matchs 这样的一个个集合，matchs[0],matchs[1],matchs[2]
	//[<a href="http://album.zhenai.com/u/1525150912" target="_blank">
	// 真善美</a> http://album.zhenai.com/u/1525150912 真善美]
	matchs := re.FindAllSubmatch(body, -1)

	result := engine.ParseResult{}
	limit := 100
	for _, m := range matchs{

		//string(m[2]) 其实这个就是用户的名字，所以可以存下来，直接用，不需要再次获取
		//所以有了下面的v2
		name := string(m[2])
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			//v1 ParseFunc:ParseUserProfile,
			//v2
			/*这边的匿名函数是在 for 结束后运行的，
			而m 的作用域 是在 for 里面的，
			所以用m[2]会出现 取到的name 是一样的情况
			 */
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseUserProfile(c, name)
			},
		})

		limit--
		if limit == 0 {
			break
		}

	}

	return result

}