package main

import "fmt"

//常量的声明用,每一行自动加一
func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const d = iota //遇到const 重置为0
	fmt.Println(d)
	const e, f = iota, iota //同一行是一样的
	fmt.Println(e, f)
	const ( //最佳
		g = iota
		k
		l
		m
	)
	fmt.Println(g, k, l, m)
	const (
		x          = iota
		x1, x2, x3 = iota, iota, iota //同一行枚举相同
		x4, x5     = iota, iota
		x6         = iota
	)
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
}
