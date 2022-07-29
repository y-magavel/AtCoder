package main

import "fmt"

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &s, &b)

	if s == "+" {
		fmt.Println(a + b)
	} else if s == "-" {
		fmt.Println(a - b)
	}
}
