package main

import "fmt"

func main() {
	a := 3
	x := func() {
		a++
	}
	x()
	fmt.Println(a)
	func() {
		a++
	}()
	fmt.Println(a)
	func(a int) {
		fmt.Println(a, "内部引用a") //符合js的传参的基本类型的知识，参数到这里还是上层的a
		a = 100
		fmt.Println(a, "a内部调用") //到这里的a就是内部的a了，不是外部的a了
	}(a)
	fmt.Println(a) //函数内对a的操作，没有改变外部的a,如果a是作为参数传入的,这是根本！
	z, y := func(a, b int) (max, min int) { //有返回值时的匿名函数
		if a > b {
			max = a
			min = b
		} else {
			max = b
			min = a
		}
		return
	}(20, 40)
	fmt.Println(z, y)
	//闭包，所有的匿名函数都是
	d := t8()   //不返回匿名函数
	fmt.Println(d)
	fmt.Println(d)
	fmt.Println(d)
	e := t9()    //返回匿名函数，匿名函数的内容一直都存在栈里面，所以暂存了
	fmt.Println(e())
	fmt.Println(e())
	fmt.Println(e())

}
func t8() int {
	var a int
	a++
	return a
}
func t9() func() int {
	var a int    //这里保证了数据的来源的唯一性，是函数体里面的数据
	return func() int {
		a++       //对a的操作，a的值会改变的，因为这里不是传的参数，是直接对a的数据的操作。
		return a
	}
}
