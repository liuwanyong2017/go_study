package main

import "fmt"

func main() {
	var a []int
	fmt.Printf("%p\n", a) //切片名本身就是一个地址， 0X0
	//append 追加数据时，内存地址，可能会发生改变
	//如果此时内存里的a的后面还有空间，指针就不会改变，否则改变
	//因为指针的地址取得就是第一个数据节点的内存地址
	a = append(a, 3, 4, 4, 4)
	fmt.Printf("内存：%p 第一个：%p\n", a, &a[0])
	//切片名本身就是一个地址， 改变了，栈区的，指针的数据

	a = append(a, 3, 4, 4, 4)

	fmt.Printf("%p\n", a) //切片名本身就是一个地址， 改变了

	//扩容
	b := make([]int, 5, 9) //预留了9个
	fmt.Println(b)
	b[3] = 7
	b = append(b, 33, 445, 2, 2)
	//小于1024 的时候，是以2倍的容量扩容的，大于的时候，是以四分之一的上一次大小扩容
	//每次扩容都是2的倍数
	fmt.Printf("len=%d,cap=%d", len(b), cap(b))
	b = append(b, 23)
	fmt.Printf("len=%d,cap=%d", len(b), cap(b))
	for i:=0;i<5500;i++ {
		b = append(b, 33, 445, 2, 2)
		fmt.Printf("len=%d,cap=%d\n", len(b), cap(b))

	}
}
