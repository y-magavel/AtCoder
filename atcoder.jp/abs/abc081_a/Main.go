package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var result int
	array := strings.Split(s, "")

	for _, v := range array {
		i, _ := strconv.Atoi(v)
		result += i
	}
	
	fmt.Println(result)
}