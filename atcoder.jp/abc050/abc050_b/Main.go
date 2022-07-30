package main

import "fmt"

func main() {
	var n int
	var t []int
	var m int
	px := make([]map[int]int, m)

	// 問題数を標準入力から取得
	fmt.Scan(&n)

	// n問の問題ごとにそれぞれ解くのにかかる時間をスライスに保持する（ex. [2, 1, 4]）
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		t = append(t, v)
	}

	// ドリンク数を標準入力から取得
	fmt.Scan(&m)

	// 各ドリンクを飲んだ際の問題を解くのにかかる時間の変化をスライスとマップで保持する（ex. [{1: 1}, {2: 3}]）
	for i := 0; i < m; i++ {
		var p, x int
		fmt.Scan(&p, &x)
		px = append(px, map[int]int{p: x})
	}

	// ドリンク数分、そのドリンクを飲んだ場合に問題を解くのにかかる時間を出力する
	for _, m := range px {
		for i, v := range m {
			result := getSolvingSpeedWithDrink(t, i, v)
			fmt.Println(result)
		}
	}

}

// あるドリンクを飲んだ際にすべての問題を解くのにかかる時間を返す
func getSolvingSpeedWithDrink(t []int, drinkNumber int, newTime int) int {
	var result int

	for i, v := range t {
		if i+1 == drinkNumber {
			result += newTime
		} else {
			result += v
		}
	}

	return result
}
