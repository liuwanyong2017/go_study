package main

import "fmt"

func main() {
	var a bool //默认值为false
	fmt.Println(a)
	a = true
	fmt.Println(a)
	//fmt.Println("%/b",a===0)
	//a=bool(1)    布尔不支持类型转换，不接受其他的类型
	b := true
	fmt.Println("b=", b)
	b = (2 == 4)
	c := (5 == 6)
	fmt.Println(b, c)
}
