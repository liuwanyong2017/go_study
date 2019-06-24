package main

import "fmt"

//不定参放最后,固定参数必须传值，不定参数，可有可无
func sum(a int, arg ...int) {
	res := 0
	for _, item := range arg {
		res += item
	}
	fmt.Println(res, a)
}
func s(a, b int) int {
	return a + b
}
func t1(args ...int) {
	t2(args[0:2]...)
	t3(args[2:4]...) //取一段的参数
	t4(args[:]...)   //所有的参数
	t4(args...)
}
func t2(args ...int) {
	fmt.Println("2", args)
}
func t3(args ...int) {
	fmt.Println("3", args)
}
func t4(args ...int) {
	fmt.Println("4", args)
}
func main() {
	sum(12, 3, 34, 3, 4545, 4)
	fmt.Println(s(2, 5))
	t1(3, 23, 8, 3, 23, 55, 43, 4)
}
