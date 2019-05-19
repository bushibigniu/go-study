package main

import (
	"fmt"
	"runtime"
	"time"
)

func main2()  {

	for i:=0;i<10 ; i++ {


		func(i int){
			for  {
				fmt.Printf("i:%d \n",i)
			}
		}(i)

		runtime.Gosched()

		//go run -race main.go //检测冲突

	}
}

func main()  {

	fmt.Println("=======begin=====")

	//var a [10]int
	//for i := 0; i<10; i++ {
	//	go func(j int){
	//		fmt.Printf("j is %d \n", j)
	//	}(i)
	//}

	for i := 0; i<1000; i++ {
		 go func(i int){
			for  {
				fmt.Printf("j is %d \n", i)
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
