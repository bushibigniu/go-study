package engine

import (
	"go-study/self/crawer/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]

		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("fetcher: error " +
				"fetcher url %s : %v", r.Url, err)
			continue
		}

		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _ , item := range parseResult.Items{
			log.Printf("get item %v", item)
		}


	}
}