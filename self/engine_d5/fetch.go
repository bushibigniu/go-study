package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//关联 city_d3 city.go

func Fetch(url string) ([]byte, error) {

	//url := "http://www.zhenai.com/zhenghun"
	resq, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	defer resq.Body.Close()
	if resq.StatusCode != http.StatusOK {
		fmt.Println("response code :", resq.StatusCode)
		return nil,
		fmt.Errorf("wrong status code: %d",
		resq.StatusCode)
	}

	e := determineEncoding(resq.Body)

	utf8Reader := transform.NewReader(resq.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	byte , err := bufio.NewReader(r).Peek(1024)

	if err != nil{
		//panic(err)
		log.Printf("Fetcher error: %v", err)
		//不要 panic ,默认返回一个utf8
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(byte, "")

	return  e
}