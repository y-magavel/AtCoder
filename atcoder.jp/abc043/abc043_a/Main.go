package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// 必要なキャンディーの個数を求める
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}
