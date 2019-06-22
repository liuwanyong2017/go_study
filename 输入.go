package main

import . "fmt"

func main() {
	//声明一个变量，以及变量的数据类型，以便于在内存中占位，
	//用scanf等待用户输出，这里阻塞进程的
	var a float64
	Println("请用户输入：")
	//fmt.Scanf("%d",&a)   //f就跟变量的声明的类型有关的，scanf的第一个参数，一定要跟变量的声明的类型一致
	//类型一致很麻烦，怎么办？
	_, _ = Scan(&a)
	Println("a=", a)
}
