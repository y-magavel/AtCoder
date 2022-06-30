package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 標準入力を受け取る
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	// 1以上n以下かつ、各桁の和がa以上b以下である、整数の合計を求める
	var result int
	for i := 0; i < n; i++ {
		// 各桁の和を求める
		num := getSumEachDigit(i + 1)

		// 各桁の和がa以上b以下の整数を結果に足していく
		if a <= num && num <= b {
			result += i + 1
		}
	}

	fmt.Println(result)
}

// 数値の各桁の総和を計算する
func getSumEachDigit(num int) int {
	// 受け取った数値を文字列に変換し、各桁に分けてスライスに詰める（ex. 201 => [2 0 1]）
	numStr := strconv.Itoa(num)
	numStrSlices := strings.Split(numStr, "")

	// 各桁の総和を求める（ex. [2 0 1] => 3）
	var result int
	for _, v := range numStrSlices {
		numSlice, _ := strconv.Atoi(v)
		result += numSlice
	}

	return result
}