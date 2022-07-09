package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var result string

	// 入力されたキー操作をひとつずつ処理する
	for _, v := range s {
		switch string(v) {
		case "0":
			result += "0"
		case "1":
			result += "1"
		case "B":
			// 文字列が空でなければ、右端の1文字を削除する
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		default:
			panic("不正なキー操作")
		}
	}

	fmt.Println(result)
}
