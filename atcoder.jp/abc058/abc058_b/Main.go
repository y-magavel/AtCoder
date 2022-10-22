package main

import "fmt"

func main() {
	var o, e string
	fmt.Scan(&o, &e)

	var pw string

	// 文字列oと文字列eから交互に文字を取り出してパスワードを復元する
	for i := 0; i < len(o+e); i++ {
		// 奇数番目の場合はoから、偶数番目の場合はeから文字を取り出す
		if i%2 == 0 {
			if i/2 < len(o) {
				pw += string(o[i/2])
			}
		} else {
			if i/2 < len(e) {
				pw += string(e[i/2])
			}
		}
	}

	fmt.Println(pw)
}
