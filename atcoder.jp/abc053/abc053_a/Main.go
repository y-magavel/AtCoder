package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch {
	case x < 1200:
		fmt.Println("ABC")
	case 1200 <= x:
		fmt.Println("ARC")
	}
}
