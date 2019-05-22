package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {

	//url := "https://www.baidu.com/"
	url := "http://www.imooc.com"
	req, err := http.NewRequest(http.MethodGet,url, nil)

	req.Header.Add("User-Agent",
		"ios")
		//"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")


	if err != nil{
		panic(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			fmt.Print("redirect:",r)
			return nil
		},
	}

	res, err := client.Do(req)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	data , err := httputil.DumpResponse(res, false)
	fmt.Printf("%s \n", data)
}
