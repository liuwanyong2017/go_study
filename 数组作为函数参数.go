package main

import "fmt"

func s3(a,b int)  {
	fmt.Printf("传入时的地址：%p\n",&b)  //这里的地址，已经是函数内的变量的地址了，不是源数据的地址了

	fmt.Printf("传入时的地址：%p\n",&a)  //这里的地址，已经是函数内的变量的地址了，不是源数据的地址了
	a,b=b,a     //存在于函数体内部的数据,执行时才有意义，但是执行完了，就会被销毁
	fmt.Println("内部",a,b)
	fmt.Printf("传入后的地址：%p\n",&a)
	fmt.Printf("传入时的地址：%p\n",&b)  //这里的地址，已经是函数内的变量的地址了，不是源数据的地址了
}
func sort(a [10]int ) [10]int{
	fmt.Printf("内部的数据包地址：%p\n",&a)  //同样也是数据拷贝了一份。
	for i:=0;i<len(a)-1 ;i++  {
		for j:=0;j< len(a)-1-i ; j++ {
			if a[j]> a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	fmt.Printf("内部的数据包地址：%p\n",&a)
	fmt.Println(a)
	return a
}
func main()  {
	a:=3
	b:=5
	s3(a,b)
	fmt.Printf("源数据的地址%p\n",&a)
	//fmt.Printf("函数的地址%p\n",&s3)
	fmt.Println("zhuhanshu",a,b,s3)

	//数组
	arr:=[...]int{33,44,5,67,45,3,24,68,92,3}
	fmt.Printf("源数据地址：%p\n",&arr)  //作为参数不影响自身，只是赋值了
	sort(arr)   //执行后就销毁了
	//如何影响源数据？？
	//函数返回值
	fmt.Println("arr=",arr)
	arr = sort(arr)
	fmt.Println("arr=",arr)

}
