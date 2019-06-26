package main

import "fmt"

func main() {
	//切片长度不固定，比数组更动态
	var a []int
	fmt.Println(a, len(a), cap(a)) //长度，容量
	a = append(a, 1, 34, 56, 6)    //添加数据
	fmt.Println(a, len(a), cap(a)) //长度，容量
	//长度是实际的数据个数，容量是再加上空的内存预留
	var b []int = []int{4, 45, 5}
	c := []int{3, 435, 5}
	fmt.Println(b, c)
	c = append(c, 56)
	fmt.Println(c)

	d := make([]int, 5, 10)                        //5个初始化的长度，10是容量
	d = append(d, 34, 45, 9, 56756, 742, 2, 2, 34) //从最后一个位置开始填值
	fmt.Println(d)
	//长度一定要小于容量，可以不设置容量，它动态添加容量

	//赋值
	s := make([]int, 5)
	s[0] = 3
	s[3] = 9
	s=append(s,45,78) //扩容的下一位添加值！
	fmt.Println(s)
}
