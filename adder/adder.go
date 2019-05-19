package main

import "fmt"

func adder() func(int) int {
	//在函数外部，自由变量
	sum := 0
	//return 的是一个闭包
	return func(v int) int { //函数体，函数体有局部变量v
		sum += v
		return sum
	}
}

//正统的的函数式编程， 递归
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}
