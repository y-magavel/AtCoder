package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)

	// a が b よりも左にある場合
	if a < b {
		// a が b の左端から w 以上離れている場合
		if b-a > w {
			fmt.Println(b - a - w)
		} else {
			fmt.Println(0)
		}
	} else {
		// a が b の右端から w 以上離れている場合
		if a-b > w {
			fmt.Println(a - b - w)
		} else {
			fmt.Println(0)
		}
	}
}
