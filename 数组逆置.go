package main

import "fmt"

func main() {
	a := [...]int{ 1,2, 4, 2, 9, 8, 3, 0, 4}
	var (
		max = a[0]
		min = a[0]
		sum int
	)
	for _, x := range a {
		sum += x
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	ping := float64(sum) / float64(len(a))
	fmt.Println(min, max, sum, ping)
	//倒置数组
	fixLen := len(a)/2
	for i:=0;i<= fixLen ;i++  {
		if i< fixLen {
			a[i] ,a[len(a)-1-i]= a[len(a)-1-i],a[i]
		}
	}
	fmt.Println(a)

	s:=0
	l:=len(a) - 1
	for s< l  {
		a[s],a[l]=a[l],a[s]
		s++
		l--
	}
	fmt.Println(a)
}
