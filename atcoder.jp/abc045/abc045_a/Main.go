package main

import "fmt"

func main() {
	var a, b, h int
	fmt.Scan(&a, &b, &h)

	// 台形の面積を求める
	area := (a + b) * h / 2

	fmt.Println(area)
}
