package main  //パッケージ

import (
	"fmt"
	"math/rand"
)	//fmt、ランダムモジュール追加

// メイン関数
func main() {
	random_value := rand.Intn(100)

	if random_value < 50 {
		goto SMALL
	} else {
		goto LARGE
	}
SMALL:
	fmt.Printf("%d is smaller than 50 \n", random_value)
	return

LARGE:
	fmt.Printf("%d is bigger than 50\n", random_value)
}