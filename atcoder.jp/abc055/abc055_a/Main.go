package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	paymentPrice := 800 * n
	discountPrice := n / 15 * 200

	fmt.Println(paymentPrice - discountPrice)
}
