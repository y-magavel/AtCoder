package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if a == "H" {
		fmt.Println(b)
	} else {
		fmt.Println(reverse(b))
	}
}

func reverse(b string) string {
	var result string

	switch b {
	case "H":
		result = "D"
	case "D":
		result = "H"
	}

	return result
}
