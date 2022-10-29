package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c, d, e, f int64
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	bigC := big.NewInt(c)
	bigD := big.NewInt(d)
	bigE := big.NewInt(e)
	bigF := big.NewInt(f)

	bigA.Mul(bigA, bigB).Mul(bigA, bigC)
	bigD.Mul(bigD, bigE).Mul(bigD, bigF)
	bigA.Sub(bigA, bigD).Mod(bigA, big.NewInt(998244353))
	fmt.Println(bigA)
}
