package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	// a*bとc*dを比較して大きい方の値を出力（同じ場合はその値を出力）
	if a*b > c*d {
		fmt.Println(a * b)
	} else if a*b < c*d {
		fmt.Println(c * d)
	} else {
		fmt.Println(a * b)
	}
}
