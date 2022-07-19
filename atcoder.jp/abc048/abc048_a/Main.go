package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力から空白区切りで読み込む設定
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	// コンテスト名を3区切りで読み込む
	var contentName [3]string // 例：AtCoder Beginner Contest
	var result string         // 例：ABC
	for i := 0; i < 3; i++ {
		sc.Scan()
		contentName[i] = sc.Text()
		result += string(contentName[i][0]) // 単語の最初の英大文字を省略名に使用する
	}

	fmt.Println(result)
}
