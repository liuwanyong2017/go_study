package main

import "fmt"

func main() {
	//float64 float32
	//双精度64 后10几位     单精度32 后7位
	a := 3.54 //自动推导的默认64
	fmt.Printf("%T\n", a)
	var (
		b         = 4.334353453453453334553
		c float32 = 3.346456456464564
	)
	fmt.Println(b, c)

}
