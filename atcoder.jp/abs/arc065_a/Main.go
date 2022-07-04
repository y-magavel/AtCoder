package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 英数字からなる文字列sを標準入力から取得する
	var s string
	fmt.Scan(&s)

	// dream、dreamer、erase、eraserのみで構成されていることを示す正規表現を作成する
	regex := regexp.MustCompile(`^(dream|dreamer|erase|eraser)*$`)

	// 文字列sが正規表現にマッチするかをチェックする
	if regex.MatchString(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}