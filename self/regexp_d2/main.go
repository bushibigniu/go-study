package main

import (
	"fmt"
	"regexp"
)

func main ()  {

	str  := `my con  wjoln a3h@www.com
	nfkw
	nw%biw qq@www.vom
	88@om.cn
	`

	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)

	reg := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)`)


	//one
	match := re.FindString(str)
	//all
	matchAll := re.FindAllString(str, -1)

	//all
	matchAll2 := reg.FindAllStringSubmatch(str, -1)

	//str to byte
	m3 := reg.FindAllSubmatch([]byte(str), -1)

	fmt.Println(match)
	fmt.Println(matchAll)
	fmt.Println(matchAll2)

	for k,v := range matchAll2{
		fmt.Println(k,v)
	}

	fmt.Println("==============")

	for _, m := range m3{
		for _, v := range m{
			fmt.Printf("%s ", v)
		}
	}

}
