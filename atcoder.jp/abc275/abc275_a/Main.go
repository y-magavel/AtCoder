package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var takahashi int
	var result int

	// 橋の数ループする
	for i := 0; i < n; i++ {
		var bridge int
		fmt.Scan(&bridge)

		// 今までの橋より高い橋もしくは、最初の橋だったらtrue
		if takahashi < bridge || i == 0 {
			takahashi = bridge
			result = i + 1
		}
	}

	fmt.Println(result)
}
