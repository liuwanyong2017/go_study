package main

import "fmt"


func main() {
	var name  = "liuliu"
	name = "wanwan"
	fmt.Println(name)
	fmt.Printf("%T", name)
	me := 56
	fmt.Println(me,name) 			//可以换行，推荐的打印方法
	fmt.Printf("%T\t",me)
	fmt.Print("\n",me,99)   //不换行
	test(me)
}
func test(num int)  {
	if num < 10 {
		fmt.Println("太小了")
	}else {
		fmt.Println("太大了")
	}
}