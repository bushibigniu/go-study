package engine

import (
	"go-study/self/15bingfa/crawer2/fetcher"
	"log"
)

type SimpleEngine struct {

}

//scheduler 实现第一步：所有worker 公用一个输入
func (e SimpleEngine) Run(seed ...Request) {

	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//brgin =====架构修改，把一下当做一个 work
			//log.Printf("fetcher url : %s", r.Url)
			//body, err := fetcher.Fetch(r.Url)
			//if err != nil {
			//	log.Printf("fetcher error fetcher url %s: %v", r.Url, err)
			//	continue
			//}
			//parseResult := r.ParserFunc(body)
		//end  work
		parseResult, err := e.worker(r)
		if err != nil {
			continue
		}

		//requests = append(requests, parseResult.Requests[0])
		//requests = append(requests, parseResult.Requests[1])
		//这是个集合， 用... 代替所有
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("get item  %v", item)
		}
	}
}

//将 parser , fetcher 封装成一个 worker
func (e SimpleEngine) worker(r Request) (ParseResult, error){
	log.Printf("fetcher url : %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error fetcher url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body),  nil
}
