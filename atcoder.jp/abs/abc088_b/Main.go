package main

import (
	"fmt"
	"sort"
)

func main() {
	// カードの枚数を標準入力から受け取る
	var n int
	fmt.Scan(&n)

	// 1枚ごとのカードの数字を標準入力からスライスで受け取る
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	// カードを数字が大きい順に並び替える
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	// AliceとBobの点差を求める
	var result int
	for index, v := range cards {
		// スライスは数字が大きい順になっているので単純に交互に取っていけばいい
		// そしてBobがAliceの点数を上回ることがないのでAliceは足し算、Bobは引き算で計算する
		if index%2 == 0 {
			result += v // Aliceは初めと偶数番目にカードを取る
		} else {
			result -= v // Bobは奇数番目にカードを取る
		}
	}

	fmt.Println(result)
}