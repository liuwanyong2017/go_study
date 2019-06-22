package main

import "fmt"

func main() {
	var a = 10
	b := 20.5
	fmt.Println(a, b)
	var (
		c = 10
		d = 20
	)
	fmt.Println(c, d)
	var (
		f = "liuliu"
		e = "haha"
	)
	fmt.Println(f, e)
	j, k, w := 45, "wan", 68.9 //最简单的多个变量声明
	fmt.Println(j, k, w)
	const me = "yong" //常量用const
	fmt.Println(me)
	const (
		h = 40
		m = 90
		i = "ooo"
	)
	fmt.Println(h, m, i)
	const l, v, n = 0, "y", 90.8 //最简单的多个命名方式
	fmt.Println(l, n, v)
}
