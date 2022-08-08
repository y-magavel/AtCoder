package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string // ex. "QWERTYASDFZXCV"
	fmt.Scan(&s)

	a := strings.Index(s, "A")     // ex. 6
	z := strings.LastIndex(s, "Z") // ex. 10

	fmt.Println(z - a + 1)
}
