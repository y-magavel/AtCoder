package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if is575(a, b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// 五七五に並び替えることができる文節か判断する
func is575(a, b, c int) bool {
	if a+b+c == 17 && (a == 5 || a == 7) && (b == 5 || b == 7) && (c == 5 || c == 7) {
		return true
	} else {
		return false
	}
}