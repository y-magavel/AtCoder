package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)

	var result = "Yes"

	// 文字列をマップにして出現回数をカウントする
	m := make(map[rune]int)
	for _, val := range w {
		m[val] = m[val] + 1
	}

	// 出現回数が奇数の文字があればNoを結果用変数に格納する
	for _, val := range m {
		if val%2 != 0 {
			result = "No"
		}
	}

	fmt.Println(result)
}
