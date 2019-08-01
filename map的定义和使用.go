package main

import "fmt"

func main() {
	//names := []string {"哈哈","u","we"}
	var x map[int]string
	fmt.Println(x,len(x))    //map[]    , 0
	//在字典中不能用cap(),只能用len()
	y := make(map[int]string,3)

	fmt.Println(y)
//key是唯一的,value是随意的，
//map内对各个key的先后顺序是随机的
//map自动扩容，不需要定容量的
	y[1] = "张三"
	fmt.Println(len(y))
	y[2] = "lisi"
	fmt.Println(len(y))
	y[3] = "ww"
	fmt.Println(len(y))
	y[4] = "haha"
	fmt.Println(len(y))

	fmt.Println(y)

	//初始化
	z := map[int]string{1:"haha",2:"dd",3:"h"}
	fmt.Println(z,z[2])

	fmt.Printf("%p\n",x)    // 内存地址为0x0   ,是系统占用，不可进行读写操作
	fmt.Printf("%p\n",y)
	fmt.Printf("%p\n",z)
}
