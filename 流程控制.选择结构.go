package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入您的分数：")
	fmt.Scan(&score)
	if score > 680 {
		fmt.Println("恭喜！清华北大在等你哦！！流弊！")
	} else {
		fmt.Printf("%d太低了！！垃圾！\n", score)
	}
	//初始化语句，需要分号去分割开条件
	if a := 650; a > 700 {
		fmt.Println("恭喜！清华北大在等你哦！！流弊！")
	} else {
		fmt.Printf("%d太低了！！垃圾！\n", a)
	}
}
