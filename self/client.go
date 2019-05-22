package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {
	url := "https://www.baidu.com"
	res, err := http.Get(url)

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()

	data, err := httputil.DumpResponse(res, false)
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)


}
