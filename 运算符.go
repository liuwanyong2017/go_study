package main

import "fmt"

func main() {
	a := 10
	a++
	fmt.Println("a=%\b", a)
	a--
	fmt.Println(a)
	b := 20
	fmt.Println(a == b)
	fmt.Println(a > 3 || b < 3)
	fmt.Println(true && false)
	fmt.Println(!(a == b))
	var c int
	c = a & b      //按位与        0000 1010
	fmt.Println(c) //     0001 0100
	//    0000 0000

	fmt.Printf("%p\n", &a) //取内存地址，%p16进制的内存地址
	fmt.Println(*&a)       //*取内存里的值
	fmt.Printf("%d", *&a)
	//优先级，自己内部的逻辑第一级，^  !
}
