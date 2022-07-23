package main

import (
	"fmt"
	"regexp"
)

func main() {
	var c string
	fmt.Scan(&c)

	// 英小文字cが母音であるかをチェックする正規表現を作成する
	regex := regexp.MustCompile(`^(a|e|i|o|u)*$`)

	// 英小文字cが母音であるかをチェックする
	if regex.MatchString(c) {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}
}
