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

/**

//
//发现站点编码
goland.org/x/net/html

//gbk 转 utf8
goland.org/x/text

 */

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



	//v3 diff begin


	bodyReader := bufio.NewReader(resq.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	data, err := ioutil.ReadAll(utf8Reader)



	//v3 old
	//e := determineEncoding(resq.Body)
	//utf8Reader := transform.NewReader(resq.Body,e.NewDecoder())


	data, err2 := ioutil.ReadAll(utf8Reader)

	//v3 diff end


	if err2 != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)


}

// 老的方式存在问题就是，peek (1024) 读出来的不完整
//func determineEncoding(r io.Reader) encoding.Encoding{ old
func determineEncoding(r *bufio.Reader) encoding.Encoding{
	byte, err := bufio.NewReader(r).Peek(1024)

	if err != nil{
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(byte,"")

	return e

}