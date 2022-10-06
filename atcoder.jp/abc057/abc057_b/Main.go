package main

import (
	"fmt"
	"math"
)

func main() {
	// n 人の学生と m 個のチェックポイント
	var n, m int
	fmt.Scan(&n, &m)

	// 学生の座標とチェックポイントの座標を標準入力から取得
	var student [][]int
	var checkPoints [][]int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		student = append(student, []int{a, b})
	}
	for i := 0; i < m; i++ {
		var c, d int
		fmt.Scan(&c, &d)
		checkPoints = append(checkPoints, []int{c, d})
	}

	// 学生ごとに一番近いチェックポイントが何番目かを出力
	for i := 0; i < n; i++ {
		fmt.Println(searchClosestCheckPoint(student[i][0], student[i][1], checkPoints))
	}
}

// 一番近いチェックポイントを探す
func searchClosestCheckPoint(a, b int, checkPoints [][]int) int {
	var min int = 0
	for i := 0; i < len(checkPoints); i++ {
		// 初回ループ時は比較対象のチェックポイントがないのでスキップする
		if i == 0 {
			continue
		}

		// 一つ前のチェックポイントより今回のチェックポイントの方が近ければ、今回のチェックポイントを一番近いチェックポイントとする
		if manhattanDistance(a, b, checkPoints[i][0], checkPoints[i][1]) < manhattanDistance(a, b, checkPoints[min][0], checkPoints[min][1]) {
			min = i
		}
	}

	return min + 1 // 1-indexed
}

// 絶対値でマンハッタン距離を求める
func manhattanDistance(a, b, c, d int) int {
	return int(math.Abs(float64(a)-float64(c)) + math.Abs(float64(b)-float64(d)))
}
