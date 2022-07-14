package main

import (
	"fmt"
)

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	// Aさん、Bさん、Cさんのカードをスライスに変換する
	var aCards, bCards, cCards []string
	aCards = stringToSlice(a)
	bCards = stringToSlice(b)
	cCards = stringToSlice(c)

	var isWin bool                                                // 勝利フラグ
	var nextTern string = "a"                                     // 次のターンの人
	var winner string                                             // 勝者を格納する変数
	var allCardsNum int = len(aCards) + len(bCards) + len(cCards) // ループ回数を格納する変数

	for i := 0; i < allCardsNum; i++ {
		switch nextTern {
		case "a":
			isWin, nextTern, aCards = execTern(aCards)
			if isWin {
				winner = "A"
			}
		case "b":
			isWin, nextTern, bCards = execTern(bCards)
			if isWin {
				winner = "B"
			}
		case "c":
			isWin, nextTern, cCards = execTern(cCards)
			if isWin {
				winner = "C"
			}
		}

		// 勝者が決まった時点でループ終了
		if winner != "" {
			break
		}
	}

	fmt.Println(winner)
}

// 文字列をスライスに変換する
func stringToSlice(s string) []string {
	var slice []string
	for _, v := range s {
		slice = append(slice, string(v))
	}
	return slice
}

// ターンを実行する
func execTern(cards []string) (isWin bool, nextTern string, newCards []string) {
	// 手持ちのカードがなければ勝利
	if len(cards) == 0 {
		return true, "", nil
	}

	nextTern = cards[0]
	newCards = removeSlice(cards)

	return false, nextTern, newCards
}

// スライスから先頭の要素を削除する
func removeSlice(slice []string) (newSlice []string) {
	if len(slice) <= 1 {
		return nil
	}

	newSlice = slice[1:]
	return newSlice
}
