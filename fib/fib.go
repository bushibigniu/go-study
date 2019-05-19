package fib

import "fmt"

//f3 = f2 + f2  fn = fn-1 + fn-2
/**

	1, 1, 2, 3, 5, 8
       a, b
		  a, b

所以是每次移动了一位，去左边数据

 */




// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main()  {
	//自己加的
	f := Fibonacci()
	fmt.Println(f())
}
