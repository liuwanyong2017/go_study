package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入您的分数：")
	fmt.Scan(&score)
	if score > 680 {
		fmt.Println("恭喜！清华北大在等你哦！！流弊！")
		if score < 700 {
			fmt.Printf("%d可以去学汽修啦！\n", score)
		} else if score < 720 {
			fmt.Printf("%d可以去学洗剪吹啦啦！\n", score)
		} else if score < 740 {
			fmt.Printf("%d可以去学挖掘机啦！\n", score)
		} else {
			fmt.Printf("%d可以去学厨师啦！\n", score)
		}
	} else if score < 200 {
		fmt.Printf("%d太低了！！蓝翔可以吗？\n", score)
	} else if score > 200 && score < 500 {
		fmt.Printf("%d还不错！可以去专科和二本啦！", score)
	} else if score > 500 && score < 600 {
		fmt.Printf("%d不错！可以去一本啦！", score)
	} else {
		fmt.Printf("%dliubi！可以去211啦！", score)
	}
	//初始化语句，需要分号去分割开条件
	if a := 650; a > 700 {
		fmt.Println("恭喜！清华北大在等你哦！！流弊！")
	} else {
		fmt.Printf("%d太低了！！垃圾！\n", a)
	}
}
