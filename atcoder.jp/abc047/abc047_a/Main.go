package main

import (
	"fmt"
)

func main() {
	var candies [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&candies[i])
	}

	// 配列candyを降順にソートする
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if candies[i] > candies[j] {
				candies[i], candies[j] = candies[j], candies[i]
			}
		}
	}

	// キャンディを2人の子供に分けられるかどうかを判定する
	if candies[0] == candies[1]+candies[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
