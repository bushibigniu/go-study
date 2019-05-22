package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {
	url := "https://www.baidu.com"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("User-Agent","ios")

	if err != nil{
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil{
		panic(err)
	}

	data , err := httputil.DumpResponse(res, true)
	fmt.Printf("%s \n", data)

}
