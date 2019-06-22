package main

import "fmt"

func main() {
	a := 43
	b := 60
	sum := (a + b) / 3
	fmt.Println(sum, float64(a+b)/3) //类型转换源于不同类之间的因果
	//类型兼容的可以互相转换，bool那里试过了，不能转换，int与float很兼容，可以转换
	//bool(0)报错
	//转换的原则，尽量向位数大的转换，越小的位数，范围越低，精度越小。
	c := 1234
	fmt.Println(float64(3.554523), float32(3.454565655675), int8(c), int64(c))
}
