package engine

import "go-study/self/crawer2/fetcher"

func Run(seed ...Request)  {
	//最终是 engine 去执行，所以写个 run

	var requests []Request

	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests)  > 0 {
		r := requests[0]
		requests = requests[1:]

		 body, err := fetcher.Fetch(r.Url)


	}


}
