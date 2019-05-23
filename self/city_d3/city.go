package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

)


func main()  {
	
	//begin  把这部分抽象出来  fetch
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


	e := determineEncoding(resq.Body)
	utf8Reader := transform.NewReader(resq.Body,e.NewDecoder())


	data, err := ioutil.ReadAll(utf8Reader)



	if err != nil{
		panic(err)
	}

	fmt.Printf("%s \n", data)
	
	//end

	printCityList(data)


}

func printCityList(content []byte)  {
	rex := `<a href="http://www.zhenai.com/[0-9a-z]+"[^>]*>[^<]+</a>`
	re := regexp.MustCompile(rex)
	//matchs := re.FindAll(content, -1)
	matchs := re.FindAllSubmatch(content, -1)

	for _,m := range matchs{
		//fmt.Printf("%s \n", m)
		fmt.Printf("city:%s , url: %s \n", m[2],m[1])
	}

	fmt.Printf("matchs found: %d\n", len(matchs))
}


func determineEncoding(r io.Reader) encoding.Encoding{
	byte, err := bufio.NewReader(r).Peek(1024)

	if err != nil{
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(byte,"")

	return e

}
