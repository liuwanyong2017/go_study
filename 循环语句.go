package main

import "fmt"

func main() {
	//初始化条件，判断条件，变化条件
	//i 块级变量
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	str := "liuliuwan"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	//元素位置，元素本身，两个返回值！
	for i, data := range str {
		fmt.Println(i, data)
	}
	for i, _ := range str {
		fmt.Println(i, "下标")
	}
	for _, data := range str {
		fmt.Println(data, "数据")
		fmt.Printf("%d\n", data)
		fmt.Printf("%c\n", data)
	}
}
