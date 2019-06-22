package main

import "fmt"

func main() {
	//complex128 默认 ，可以 64
	a := 2.1 + 4.3i
	var b complex64 = 3 + 4.3 //不加i，默认为0

	fmt.Println(a, b)
	fmt.Printf("%T,%T\n", a, b)
	//实部为加号前的不带i的总结果，虚部为带i的常量
	fmt.Println(real(a), real(b), imag(a), imag(b))
}
