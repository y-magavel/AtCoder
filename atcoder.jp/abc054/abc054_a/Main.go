package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// カードの強さ：弱 2 < 3 < 4 < 5 < 6 < 7 < 8 < 9 < 10 < 11 < 12 < 13 < 1 強
	if a == 1 {
		a = 14
	}
	if b == 1 {
		b = 14
	}

	if b < a {
		fmt.Println("Alice")
	} else if a < b {
		fmt.Println("Bob")
	} else if a == b {
		fmt.Println("Draw")
	}
}
