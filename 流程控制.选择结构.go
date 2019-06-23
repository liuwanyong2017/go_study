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
	switch { //不建议嵌套，适合离散型的
	case score > 740:
		fmt.Printf("%d可以去学厨师啦！\n", score)
	case score > 720:
		fmt.Printf("%d可以去学挖掘机啦！\n", score)
	case score > 700:
		fmt.Printf("%d可以去学洗剪吹啦！\n", score)
		fallthrough //匹配后，继续匹配下面的，默认为break形式
	case score > 680, score > 650, score > 640:
		fmt.Printf("%d可以去985啦！\n", score)
	default:
		fmt.Printf("%d你流弊啊！", score)
	}
	switch s1 := 650; score { //如果score没有值，就用默认值
	case 650:
		fmt.Printf("%d默认值测试！", s1)
	case 640, 630, 620:
		fmt.Printf("%d玄学啊！", s1)
	default:
		fmt.Println("不知道你考得啥意思！")
	}
}
