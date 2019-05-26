package engine

import (
	"go-study/self/15bingfa/crawer/fetcher"
	"log"
)

func Run(seed ...Request) {

	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetcher url : %s", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher error fetcher url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		//requests = append(requests, parseResult.Requests[0])
		//requests = append(requests, parseResult.Requests[1])
		//这是个集合， 用... 代替所有
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("get item  %v", item)
		}
	}

}
