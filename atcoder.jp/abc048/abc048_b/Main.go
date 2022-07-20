package main

import (
	"fmt"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	var dividedA int = a / x
	var dividedB int = b / x

	// a以上b以下の整数のうち、xで割り切れるものの個数はそれぞれの整数をxで割った数の差を求めることでわかる
	var result int = dividedB - dividedA
	if a == 0 {
		result++ // 0もxで割り切れる判定なので、aが0の場合は計算結果に+1
	} else if a%x == 0 {
		result++ // aがxで割り切れる場合は計算結果が1少なくなるので+1
	}

	fmt.Println(result)
}
