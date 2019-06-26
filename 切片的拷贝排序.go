package main

import "fmt"

func main() {
	a := []int{23, 235, 34, 234, 74}
	b := make([]int, 5)
	copy(b, a) //深拷贝，地址不一样的
	fmt.Println(b)
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b) //地址不相等
	b[1] = 2
	//fmt.Println(b)
	a = append(a, 3, 43, 54)
	fmt.Println(a, b)
	copy(a, b)
	fmt.Println(a, b)
	copy(b, a[4:]) //从下表0开始覆盖原数据。
	fmt.Println(b, a)
	//排序
	for i := 0; i < len(a)-1; i++ {
		for j:=0;j<len(a)-1-i ;j++  {
			if a[j]>a[j+1] {
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	fmt.Println(a)
}
