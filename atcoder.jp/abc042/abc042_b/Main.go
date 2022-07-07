package main

import "fmt"

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	// 文字列をスライスに詰める
	var slice []string
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		slice = append(slice, s)
	}

	// スライスを辞書順にソートする
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	// スライスの中身を結合して出力する
	var result string
	for i := 0; i < n; i++ {
		result += slice[i]
	}
	fmt.Println(result)
}