package main

import "fmt"

func main() {
	//类型别名的意思就是不用官方的数据类型的那个名称，其实数据类型还是那个，只不过名称变了
	//前面遇到过byte 的数据类型的数据Printf查的时候是int8，也就是byte就是int8的别名
	type big int64
	var a big = 1231242
	fmt.Printf("%T\n", a)
	type (
		f float64
		F float32
	)
	var b f = 3.343434343
	var c F = 3.43
	fmt.Printf("%T,%T", b, c)
}
