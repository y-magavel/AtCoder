package main
 
import (
	"fmt"
	"math"
)
 
func main() {
	var n int
	fmt.Scan(&n)
 
	var result string = "Yes"
	var currentTime int = 0
	var currentX int = 0
	var currentY int = 0
 
	// 旅行プランn回分検証する（途中で実行不可能のプランがあった場合はその時点でループを抜ける）
	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		t -= currentTime                              // 前回時刻から残り時間を求める
		xDist := int(math.Abs(float64(x - currentX))) // 前回座標から残りのx距離を求める
		yDist := int(math.Abs(float64(y - currentY))) // 前回座標から残りのy距離を求める
 
		// 時刻tと、x,yの合計が偶数または奇数同士でない場合
		if !((t%2 == 0 && (xDist+yDist)%2 == 0) || (t%2 != 0 && (xDist+yDist)%2 != 0)) {
			result = "No"
			break
		}
 
		// 時刻tよりxDist,yDistの合計が大きい場合
		if t < xDist+yDist {
			result = "No"
			break
		}
 
		currentTime += t // 現在時刻にtを足す
		currentX = x    // 現在座標にxをセットする
		currentY = y    // 現在座標にyをセットする
	}
 
	fmt.Println(result)
}