package main

import "fmt"

func main()  {
	a:=[]int{3,4,44,5,66,67}
	b:=a[2:5:6]  //开始位置，包含，结束位，不包含，最后的是容量6-2,不能大于目标的容量值
	fmt.Println(b,len(b),cap(b))
	//c:=s[:] //全部截取    c:=s
	c:=a[0:3]
	d:=a[4:]
	fmt.Println(c,d)
	fmt.Printf("%p\n",c)
	fmt.Printf("%p\n",a)  //截取后的切片还是原始切片中的数据。地址一样的
	c[2]=99
	fmt.Println(c,a)//都变了

}
