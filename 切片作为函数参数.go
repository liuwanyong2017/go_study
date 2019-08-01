package main

import "fmt"
func x(m []int ){
	fmt.Println('m',m)   //形参是实参的指针
	m[0] = 234
}
func y(m []int ) []int  {
	m = append(m,5,7,9)     //这里的切片地址可能会发生改变，append
	return m
}
func main() {
	a:=[]int{1,34,5,65,6}
	fmt.Println('a',a)
	x(a)					//这里的形参传递的是切片的地址，所以跟JS的一样，会对传递的切片进行操作
	fmt.Println('a',a)
	fmt.Println(y(a),a)    //这里可能会不一样的。
}
//切片在内存中，如果没有变量指向它，会被销毁回收