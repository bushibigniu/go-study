package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	_ "io"
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
