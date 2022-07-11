package main

import "fmt"

func main() {
	// 予定宿泊数n、最初のk泊までの1泊あたりの料金x円、それ以降の1泊あたりの料金y円
	var n, k, x, y int
	fmt.Scan(&n, &k, &x, &y)

	var totalPrice int

	for i := 0; i < n; i++ {
		if i < k {
			totalPrice += x
		} else {
			totalPrice += y
		}
	}

	fmt.Println(totalPrice)
}
