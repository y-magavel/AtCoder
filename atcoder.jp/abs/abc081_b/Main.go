package main

import (
	"fmt"
)

func main() {
	// 操作回数の結果
	var result int = 0

	// 整数の個数を受け取る
	var n int
	fmt.Scan(&n)

	// 整数をスライスに詰める
	var a = make([]int, n)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		a[i] = tmp
	}

	// スライスの整数がすべて偶数の場合
	for isAllEven(a) {
		// 2で割った値に置き換える
		for i := 0; i < n; i++ {
			a[i] = a[i] / 2
		}
		// 操作回数を記録する
		result++
	}

	fmt.Println(result)
}

// スライスの要素すべてが偶数か判定する
func isAllEven(slice []int) bool {
	var result bool = true
	for _, v := range slice {
		if v%2 != 0 {
			result = false
		}
	}

	return result
}