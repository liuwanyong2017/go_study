package main

import "fmt"

func main() {
	a, b, c := 2, "ef", 64
	a, c = c, a    //这里是精髓
	fmt.Println(a, b, c)
	//var d = a
	//a = c
	//c = d赋值，麻烦
	var x,_  = 50,12   //匿名变量，下划线，被丢弃的数据不被处理
	fmt.Println(x)

	e,f,d:=get()			//匿名变量适合多个赋值中，除去某个值，因为变量一旦声明，就必须要被引用，所以匿名变量解决了。
	fmt.Println(e,f,d)
	_,d,f=get()
	fmt.Println(d,f)
}

func get()(a,b,c int)  {
	return 3,5,7
}
