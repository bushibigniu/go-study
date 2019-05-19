package main

import (
	"fmt"
	"time"
)

//1.channel 作为参数使用
func work(c chan int)  {
	for  {
		n := <- c
		fmt.Println(n)
	}
}

//2.return channel
//func createWork()  chan<- int{  //定义chan 只能发，不能收，定义的变量也要保持一致
func createWork()  chan<- int{
	c := make(chan int)

	//防止死锁，自己要发
	go func(){
		fmt.Println("crete:",<-c)
	}()
	return c
}

func bufferChannel()  {
	ch  := make(chan int, 10)
	ch <- 1
}

func closeChannel()  {
	ch  := make(chan int, 10)

	go func(){
		//第一种：加判断
		//for {
		//	n ,ok := <- ch
		//	//判断 ch 是否为空
		//	if !ok{
		//		break
		//	}
		//	fmt.Println(n)
		//}

		//第二种：range
		for n := range ch{
			fmt.Println(n)
		}
	}()

	ch <-124
	ch <-145
	ch <-14
	ch <-12
	ch <-12

	//一定是发送方close
	close(ch)

}

func main()  {

	closeChannel()

	c := make(chan int)

	//go func(){
	//	for  {
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//
	//}()

	//对于 channel 来说，可以一直输出，但不能一直 输入，因为它是阻塞型的
	go work(c)

	time.Sleep(time.Millisecond)
	c <- 1
	c <- 2

	var ch [10]chan int
	//var chslice []chan int
	for i:=0;i<10 ;i++  {

		ch[i] = make(chan int)
		//chslice[i] = make(chan int)

		go work(ch[i])
		//go work(chslice[i])
	}

	for i:=0;i<10 ;i++  {
		ch[i] <- 'a'+i
		//chslice[i] <- 'A'+i
	}


	//跟方法保持一致
	var ch2 [10]chan<- int
	for i:=0;i<10 ;i++  {
		ch2[i] = createWork()
	}

	for i:=0;i<10 ;i++  {
		ch2[i]<- i+1000
	}



	time.Sleep(time.Millisecond)
}

/**
总结：

1.channel 阻塞型
	对于 channel 来说，可以一直输出，但不能一直 输入，因为它是阻塞型的


2.channel 可以作为参数使用

3.func 可以 return channel

4. channel 可以close
   一定是发送方执行 close
	close 之后判断为空方式：

	//第一种：加判断
		//for {
		//	n ,ok := <- ch
		//	//判断 ch 是否为空
		//	if !ok{
		//		break
		//	}
		//	fmt.Println(n)
		//}

		//第二种：range
		for n := range ch{
			fmt.Println(n)
		}

5. //func createWork()  chan<- int{  //定义chan 只能发，不能收，定义的变量也要保持一致

*/
