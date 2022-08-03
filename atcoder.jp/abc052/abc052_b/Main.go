package main

import "fmt"

func main() {
	var n int    // ex. n = 5
	var s string // ex. s = "IIDID"
	fmt.Scan(&n, &s)

	var x int = 0
	var highScore int = 0

	// Iならばxの値を1増やし、Dならばxの値を1減らして、操作の途中かどうかに関わらずxがとる値の最大値を求める
	for _, v := range s {
		if string(v) == "I" {
			x++
		}
		if string(v) == "D" {
			x--
		}
		if highScore < x {
			highScore = x
		}
	}

	fmt.Println(highScore)
}
