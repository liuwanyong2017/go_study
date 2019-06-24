package main

import "fmt"

func main() {
	//初始化条件，判断条件，变化条件
	//i 块级变量
	sum := 0
	for i := 1; i < 10; i++ {
		sum = sum + i
		fmt.Println(i, sum)
	}
	str := "liuliuwan"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	//元素位置，元素本身，两个返回值！
	for i, data := range str {
		fmt.Println(i, data)
	}
	for i, _ := range str {
		fmt.Println(i, "下标")
	}
	for _, data := range str {
		fmt.Println(data, "数据")
		fmt.Printf("%d\n", data)
		fmt.Printf("%c\n", data)
	}
	//嵌套
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}
	//百鸡问题
	count := 0
	count1 := 0
	for man := 0; man <= 20; man++ {
		for woman := 0; woman <= 33; woman++ {
			child := 100 - man - woman
			count++
			if child%3 == 0 && (man*5+woman*3+child/3 == 100) {
				fmt.Printf("man:%d,woman:%d,child:%d\n", man, woman, child)
			}
		}
	}
	for man := 0; man <= 20; man++ {
		for woman := 0; woman <= 33; woman++ {
			for child := 0; child <= 100; child += 3 {
				count1++
				if child+woman+man == 100 && (child/3+woman*3+man*5 == 100) {
					fmt.Printf("man:%d,woman:%d,child:%d\n", man, woman, child)
				}
			}
		}
	}
	fmt.Printf("count=%d,count1=%d", count, count1)
}
