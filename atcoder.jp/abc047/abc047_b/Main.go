package main

import "fmt"

func main() {
	// 長方形の右上の座標点w,hと、長方形の中の座標点の個数nを標準入力から取得する
	var w, h, n int
	fmt.Scan(&w, &h, &n)

	//  n個の座標点を標準入力から取得する
	var points [][3]int
	for i := 0; i < n; i++ {
		var x, y, d int
		fmt.Scan(&x, &y, &d)
		points = append(points, [3]int{x, y, d})
	}

	// 長方形の白い部分の4隅の座標点を求める
	var left, right, top, bottom int = 0, w, h, 0
	for _, point := range points {
		switch point[2] {
		case 1: // 左側を黒く塗りつぶす
			left = max(left, point[0])
		case 2: // 右側を黒く塗りつぶす
			right = min(right, point[0])
		case 3: // 下側を黒く塗りつぶす
			bottom = max(bottom, point[1])
		case 4: // 上側を黒く塗りつぶす
			top = min(top, point[1])
		}
	}

	// 長方形の面積を求める
	var whiteArea int = (right - left) * (top - bottom)
	if whiteArea < 0 || right < left || top < bottom {
		whiteArea = 0
	}

	fmt.Println(whiteArea)
}

// 2つの整数のうち値が小さい方を返す
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 2つの整数のうち値が大きい方を返す
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
