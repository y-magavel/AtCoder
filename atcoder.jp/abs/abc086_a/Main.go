package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var result string
	if (a*b)%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}

	fmt.Println(result)
}