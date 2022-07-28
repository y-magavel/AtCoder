package main

import "fmt"

func main() {
	// 縦hピクセル、横wピクセルの画像を標準入力から読み込む
	var h, w int
	fmt.Scan(&h, &w)

	// 画像を1行ずつのピクセル（文字列）ごとにスライスに詰める
	var pixels []string
	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		pixels = append(pixels, s)
	}

	// 画像を縦方向に2倍にする
	var doubledHeightPixels []string
	for i := 0; i < h; i++ {
		doubledHeightPixels = append(doubledHeightPixels, pixels[i])
		doubledHeightPixels = append(doubledHeightPixels, pixels[i])
	}

	// 縦方向に2倍にした画像を出力する
	for _, v := range doubledHeightPixels {
		fmt.Println(v)
	}

}
