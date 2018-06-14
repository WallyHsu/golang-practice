package main

import "fmt"

func main() {
	//練習題1
	var score [5]int
	var total int

	score[0] = 59
	score[1] = 70
	score[2] = 85
	score[3] = 30
	score[4] = 100

	for _, v := range score {
		total = total + v
	}
	//平均
	average := float32(total) / float32(len(score))
	fmt.Println("各個學生成績：", score)
	fmt.Println("平均分數：", average)

	//練習2
	var min int
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min = x[0]
	for _, y := range x {
		if y < min {
			min = y
		}
	}
	fmt.Println("最小數字：", min)
}
