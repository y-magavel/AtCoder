package main

import "fmt"

func main() {
	// 標準入力から紙幣の枚数を受け取る
	var n int
	fmt.Scan(&n)

	// 標準入力から合計金額を受け取る
	var total int
	fmt.Scan(&total)

	// 10000円札がx枚、5000円札がy枚、1000円札がz枚であるとする
	var x, y, z int = -1, -1, -1

	// 合計金額を紙幣の枚数で表現できるかどうかを判定する
	for a := 0; a <= n; a++ {
		for b := 0; b <= n-a; b++ {
			for c := 0; c <= n-a-b; c++ {
				if a+b+c == n && a*10000+b*5000+c*1000 == total {
					x, y, z = a, b, c
					break
				}
			}
		}
	}

	// 結果を出力する
	fmt.Println(x, y, z)
}