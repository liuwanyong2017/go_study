package main

import "fmt"

func main() {
	a := 45
	b := 'a'
	c := "abs"
	d := 4.344343433
	//%T,%c ,%s,%d,f
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)
	fmt.Printf("%d,%c,%s,%f\n", a, b, c, d)
	fmt.Printf("%.3f\n", d) //省略小数点后几位
	//自动格式匹配 %v
	fmt.Printf("%v,%v,%v,%v", a, b, c, d)
}
