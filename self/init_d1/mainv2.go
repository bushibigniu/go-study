package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
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

	//v2 diff begin
	//如果原来站点是 gbk 可以通过这种方式写死转换至utf8
	//存在问题：就是不智能
	utf8Reader := transform.NewReader(resq.Body,
		simplifiedchinese.GBK.NewDecoder())

	data, err := ioutil.ReadAll(utf8Reader)

	//v2 diff end


	if err != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)

}