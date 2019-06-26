package main

import "fmt"

func main()  {
	var  a [3][4]int
	fmt.Println(a)
	//3行，4个长度的子数组
	fmt.Println(len(a),a[0])  //3行
	//4列
	a[0][2]=2   //0行2列
	fmt.Println(a,a[0],len(a[0]))
	//外层是行，内层是列
	for _,x:= range a   {
		for _,j :=range x  {
			fmt.Println(j)
		}
	}
	b:= [2][3]int{{3,2,5},{568,7885,3}}
	fmt.Println(b)
}
