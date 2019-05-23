package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//v2 加字符版
func main()  {
	url := "http://www.zhenai.com/zhenghun"
	resq, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	defer resq.Body.Close()

	if resq.StatusCode != http.StatusOK {
		fmt.Println("response code :", resq.StatusCode)
		return
	}

  	data, err := ioutil.ReadAll(resq.Body)
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)

}
