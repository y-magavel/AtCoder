package main

import "fmt"

func main() {
	// 500円玉がa枚、100円玉がb枚、50円玉がc枚、合計金額がx円
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	// 何通りあるかの結果用変数
	var result int

	// 小さい小銭から1枚ずつ増やしていく
	for i500 := 0; i500 <= a; i500++ {
		for i100 := 0; i100 <= b; i100++ {
			for i50 := 0; i50 <= c; i50++ {
				if 500*i500+100*i100+50*i50 == x {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}
