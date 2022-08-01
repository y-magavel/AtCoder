package main

import "fmt"

func main() {
	var k, s int
	fmt.Scan(&k, &s)

	var result int

	// x+y+z=k, x+y+z=sを満たすx,y,zへの値の割り当てが何通りあるかを求める
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - x - y // xとyが決まった時点でzが決まる
			if z <= k && z >= 0 {
				result++
			}
		}
	}

	fmt.Println(result)
}
