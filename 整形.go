package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 10
	a = -10
	var (
		b uint = 20
		c int  = -20
		//d uint = -2   uint 无符号整形，int有符号
		//占4或者8字节
	)
	fmt.Println(a, b, "c=%\b", c)
	fmt.Printf("%T\n", b)
	fmt.Println(unsafe.Sizeof(b))
	var d int32 = 20
	fmt.Println(unsafe.Sizeof(d)) //int后面可以跟多少位的，16，8，32，64,一般64位，因为范围大，不过占的内存也大两个字节
}
