package main

import (
	"fmt"
)

func main() {
	// 餅の枚数を受け取る
	var n int
	fmt.Scan(&n)

	// 餅1枚ごとの直径をスライスで受け取る
	mochiList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mochiList[i])
	}

	// 同じ直径の餅は1つにする（鏡餅が何段になるかは直径が違う餅がいくつあるかで決まるため）
	uniqMochiList := removeDuplicate(mochiList)

	fmt.Println(len(uniqMochiList))
}

// スライスから重複している要素を取り除く
func removeDuplicate(slice []int) []int {
	m := make(map[int]bool)
	uniqSlice := make([]int, 0)

	for _, v := range slice {
		if !m[v] {
			m[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}

	return uniqSlice
}