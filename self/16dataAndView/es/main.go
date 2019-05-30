package main

import (
	"github.com/elastic"
)

func main()  {

	client,err := elastic.NewClient()

	if err != nil {
		panic(err)
	}

	//todo
	client.Index().Type("a").Id("b")


}
