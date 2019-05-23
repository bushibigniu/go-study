package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
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



	//v3 diff begin

	e := determineEncoding(resq.Body)
	utf8Reader := transform.NewReader(resq.Body,e.NewDecoder())


	data, err := ioutil.ReadAll(utf8Reader)

	//v3 diff end


	if err != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)


}
func determineEncoding(r io.Reader) encoding.Encoding{
	byte, err := bufio.NewReader(r).Peek(1024)

	if err != nil{
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(byte,"")

	return e

}