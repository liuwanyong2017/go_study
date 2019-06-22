package main

import "fmt"

func main() {
	//单引号是字符，双引号是字符串
	a := 'a'
	var b byte = 'a' //byte 是uint8的别名，
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(a, b) //计算机只能识别数字，ASCll码 0到255
	//a ,97 往后 26， 大写的A是65，往后26个
	//0到9，从48开始

	//%c 占位符，输出一个字符，而不是数据类型
	fmt.Printf("%c\n", b)
	fmt.Printf("%c\n", 97) //97是'a'的ascll码

	var d byte = '0'
	fmt.Printf("%c\n", d)
	fmt.Println(d)
	// \n \t
	// \0字符串末尾，结束的标志，ascll 为 0 。
	fmt.Println('\n')
	start()
	qubie() //字符与字符串的区别
}
func start() {
	str := "haha"
	fmt.Println(str)
	fmt.Printf("%T\n", str)
	str = "a" // 'a'加上'\0'结束的字符
	//x:='a'  //字符
	m := "刘刘哈"                    //1汉字占3个字符
	fmt.Println(len(str), len(m)) //不包含'\0'的，长度算法

	n := "有双精度"
	fmt.Println(m + n) //拼接，就是加法
}
func qubie() {
	a := 'a' //ascll码是值，一般一个字符组成，除了\n ,\t之类
	b := "a" //一个多个字符组成，最后加个 \0，
	fmt.Println(a, b)
	//字符串的下标属性，可以拿到单个下标那里的字符
	fmt.Println(("haha")[2])                      //'h'的ascll码
	fmt.Printf("%c,%c", ("haha")[0], ("haha")[1]) //字符的值
}
