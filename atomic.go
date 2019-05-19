package main

import (
	"fmt"
	"sync"
	"time"
)


//这种方式 是传统的同步机制，
//实现了 原子锁，但还是使用共享内容内存的方式
//而不是通过通信来实现，所以还是建议使用 channel 来实现



//type atomicInt struct {
//	value int
//	lock  sync.Mutex
//}

type atomicInt int


func (a *atomicInt) increment() {
	//fmt.Println("safe increment")
	//func() {
	//	a.lock.Lock()
	//	defer a.lock.Unlock()
	//
	//	a.value++
	//}()

	*a++
}

func (a *atomicInt) get() int {
	//a.lock.Lock()
	//defer a.lock.Unlock()
	//
	//return a.value
	return int(*a)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
