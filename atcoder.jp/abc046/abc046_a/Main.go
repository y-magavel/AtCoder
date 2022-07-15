package main

import "fmt"

func main() {
	var paintColorsArray [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&paintColorsArray[i])
	}

	// 配列から重複を除去
	uniqPaintColorsSlice := removeDuplicate(paintColorsArray)

	fmt.Println(len(uniqPaintColorsSlice))
}

// 配列から重複している要素を取り除きスライスを返す
func removeDuplicate(array [3]int) []int {
	m := make(map[int]bool)
	uniqSlice := make([]int, 0)

	for _, v := range array {
		if !m[v] {
			m[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}

	return uniqSlice
}
