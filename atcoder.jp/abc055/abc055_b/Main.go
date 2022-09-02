package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	power := big.NewInt(1)
	for i := 0; i < n; i++ {
		power = power.Mul(power, big.NewInt(int64(i+1)))
	}

	result := power.Mod(power, big.NewInt(int64(math.Pow10(9)+7)))

	fmt.Println(result)
}
