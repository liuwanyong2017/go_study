package main

import "fmt"

func main() {
	a := make(map[string][3]int)
	//key支持 == ！ = 一般建议为一般类型，切片，函数，切片的结构类型都会报错
	//  map[[]string]int  报错
	//map中数据是无序的
	a["小米"] = [3]int{34,67,9}
	a["小米1"] = [3]int{345,67,96}
	a["小米2"] = [3]int{34,657,29}

	fmt.Println(a)
	for k,v := range a {
		fmt.Println(k,v)
	}

	b := make(map[int]string)
	b[1] = "刘"
	b[2] = "操"
	b[3] = "sun"
	fmt.Println(b[2],b[4])
	//b[4] 为默认值空字符串，找不到就返回一个类型的默认值。

	value,key := b[1]
	fmt.Println(value,key)     //判断是否存在键值对

	delete(b,2)
	fmt.Println(b)     //删除操作，删除的是key，删除不存在的key，不会报错
}
