package main

import "fmt"

func main() {
	count:=0
	a := [...]int{43, 4, 6, 5, 43, 2678, 2, 323, 343, 34}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			count++
			if a[j]>a[j+1] {
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	//冒泡排序的单个循环就是找到已知的最大值，一个一个地确定，第一名，第二名，，，，
	fmt.Println(a,count)
}
