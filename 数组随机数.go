package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//随机数
	a:=[...]int{}
	fmt.Println(a)
	//随机数种子
	rand.Seed(1)  //这里不变，生成的随机数不变的，它是个参数，
	fmt.Println(rand.Int())  //生成比较大的随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())   //很大的随机数
	fmt.Println(rand.Intn(10))   //0-9
	fmt.Printf("%.3f\n",rand.Float64())
	//双色球  红33 个，选6个 篮球16个，选一个 ，可以跟红球一样
	b:=[7]int{6:rand.Intn(16)+1}
	for i:=0;i<6 ;i++  {
		c:=rand.Intn(33)+1
			for j:=0;j<i;j++ {    //比较的是遍历顺序生成的前几项，所以只对比已生成的前几项
				if b[j] == c{      //一旦发现有了
					c = rand.Intn(33)+ 1    //重新生成
					j=-1            //初始化j，它会自增成0；
				}
			}
			b[i]=c
	}
	fmt.Println(b)

}
