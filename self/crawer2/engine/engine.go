package engine

import (
	"go-study/self/crawer2/fetcher"
	"log"
)

/**
	engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	}
	这个是传进来的参数，是一个struct，也就是下面使用的r,
	r.Url 值就是  url
	r.ParserFunc 就是 parser.ParseCityList
	所以 会调起来 ParseCityList 方法，获取城市列表

 */

/**
	s 第一次获取的是：http://www.zhenai.com/zhenghun
	s 这个地址结果，然后将获取到的城市列表继续，往下访问，所以会看到一直在请求
	s 但是 在 http://www.zhenai.com/zhenghun 这个地址 获取的结果，使用 了 NilParse，所以获取的地址就没有往下请求了
*/

func Run(seed ...Request) {
	//最终是 engine 去执行，所以写个 run

	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		//拿第一个 处理
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetcher url : %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher: error "+
				"fetcher url %s : %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		//Requests 这是集合，每一个都要请求，所以需要一直写下去， 用 ... 代替了所有
		//requests = append(requests,
		//	parseResult.Requests[0])
		//requests = append(requests,
		//	parseResult.Requests[1])
		requests = append(requests,
			parseResult.Requests...)

		for _, item := range parseResult.Items{
			log.Printf("get item %v", item)
		}

	}

}
