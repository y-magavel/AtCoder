package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var result int

	// 一列に並んだn個のボールをk色のペンキで、隣同士が同じ色にならない塗り方が何通りあるか求める
	for i := 0; i < n; i++ {
		if i == 0 {
			result = k
		} else {
			result = result * (k - 1)
		}
	}

	fmt.Println(result)
}
