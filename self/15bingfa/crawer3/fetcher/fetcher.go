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
	"time"
)
//time.Tick 每100ms 发送一个 request ，防止反爬虫

var rateLimit = time.Tick(time.Millisecond*1000)
//var timeAfter = time.After(time.Second)


func Fetch(url string) ([]byte, error) {

	<- rateLimit
	//v1 resp, err 直接请求
	//resp, err := http.Get(url)

	//v2 resp, err 因为访问用户profile 403，所以加了ua
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	//User-Agent: Mozilla/5.0
	request.Header.Add("User-Agent", "Firefox/11.0")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("http get err %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//编码问题 gdk utf8
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {

	byte, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(byte, "")
	return e
}
