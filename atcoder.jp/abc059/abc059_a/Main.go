package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)

	var result = strings.ToUpper(string(s1[0]) + string(s2[0]) + string(s3[0]))

	fmt.Println(result)
}
