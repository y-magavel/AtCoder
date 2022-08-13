package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var imageA, imageB []string

	// 画像Aを読み込む
	for i := 0; i < n; i++ {
		var row string
		fmt.Scan(&row)
		imageA = append(imageA, row)
	}

	// 画像Bを読み込む
	for i := 0; i < m; i++ {
		var row string
		fmt.Scan(&row)
		imageB = append(imageB, row)
	}

	// 画像の平行移動のみ許されるとき、テンプレート画像Bが画像Aの中に含まれているかを判定する
	var result bool = false
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			ans, _ := check(i, j, imageA, imageB)
			if ans {
				result = true
				break
			}
		}
	}

	// 結果を出力する
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func check(startI, startJ int, a, b []string) (bool, error) {
	var result bool = true
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if a[i+startI][j+startJ] != b[i][j] {
				result = false
				break
			}
		}
	}

	return result, nil
}
