package parser

import (
	"go-study/self/15bingfa/crawer6/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1997037559" target="_blank">安静</a>
const regexpUser2  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

//next page http://www.zhenai.com/zhenghun/shanghai/2
const rePage = `href="h(ttp://www.zhenai.com/zhenghun/[^"]+)"`


func ParsePageUser(body []byte)  engine.ParseResult{

	re := regexp.MustCompile(regexpUser2)

	//matchs 这样的一个个集合，matchs[0],matchs[1],matchs[2]
	//[<a href="http://album.zhenai.com/u/1525150912" target="_blank">
	// 真善美</a> http://album.zhenai.com/u/1525150912 真善美]
	matchs := re.FindAllSubmatch(body, -1)

	result := engine.ParseResult{}
	limit := 8
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

	rePageUser := regexp.MustCompile(rePage)
	matchs = rePageUser.FindAllSubmatch(body, -1)
	for _, m := range matchs{
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParserFunc:ParsePageUser,
		})
	}




	return result

}