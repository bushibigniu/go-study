package fetcher

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

func Fetch(url string) ([]byte, error){

	//1.request
	resq, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resq.Body.Close()

	if resq.StatusCode != http.StatusOK{
		return nil,
			fmt.Errorf("wrong status code: %d",
				resq.StatusCode)
	}

	//2.encode
	e := determineEncoding(resq.Body)

	utf8Reader := transform.NewReader(resq.Body, e.NewDecoder())

	//3.read content
	return ioutil.ReadAll(utf8Reader)
}

//对获取编码 封装
func determineEncoding(r io.Reader)  encoding.Encoding{
	byte, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		//error 返回去默认 utf8,在打个log
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(byte,"")

	return e
}
