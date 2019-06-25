package main

import "fmt"

func tt(a int) int {
	//tt(a)   //自己调自己，  防止无限次，内存不够用啊
	//所以有终止调用的操作
	if a > 20 {
		fmt.Println(a,"正好的！")
		return a
	}else {
		fmt.Println(a,"太小了！")
		return tt(a+1)
	}
}
func tt1(a int)int  {  //递归
	if a == 1 {
		return a
	}else {
		return a + tt1(a-1)
	}
}
func main() {
	//tt(2)   内存溢出啊
	tt(25)
	tt(18)
	fmt.Println(tt1(100))
}
