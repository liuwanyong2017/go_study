package main

import "fmt"

func main() {
	var a [10]  int              //长度，类型，初始值为0,内存里开辟了空间，长度为常量，十个内存空间是连续的
	fmt.Println(a, a[1], len(a)) //a本身的值就是数组的值，而不是指针,顺序存储，链式结构
	fmt.Printf("%p\n", &a)       //内存地址
	for _, data := range a {
		fmt.Printf("%p\n", &data)
	}
	//赋值
	a[0] = 3
	//a[10] = 100  超长度会报错的
	fmt.Println(a[0], a)
	for i := 0; i < len(a); i++ {
		a[i] += i
	}
	fmt.Println(a)
	var (
		b [10] string //空字符
		c [10]float64 //默认值为0
	)
	fmt.Println(b, b[0], c)

	//初始化
	//全部初始化
	var arr [5]int = [5]int{1, 2, 4, 5, 3}
	//部分初始化
	arr1 := [5]int{2, 33, 434, 4, 2}
	arr2 := [5]int{2, 33, 4,}
	arr3 := [5]int{2: 5, 4: 10}
	//长度自动初始化
	arr4 := [...]int{23, 44, 4, 5}
	arr5 := [...]int{23: 44, 4,5}  //这时候是26的长度，最后为4，5。


	fmt.Println(arr, arr1, arr2, arr3, arr4,arr5)
	x:=arr1
	fmt.Println(x,arr1,x==arr1) //指向的是一个内存??
	fmt.Printf("%T\n",x)
	x[0]=30
	fmt.Println(x,arr1,x==arr1)   //不是一个内存啊！！
	fmt.Printf("%p\n",&x)
	fmt.Printf("%p\n",&arr1)  //地址为第一个子节点的储存地址
	for i,_:= range arr1   {
		fmt.Printf("%p\n",&arr1[i])  //都是相邻的储存地址码
	}
}
