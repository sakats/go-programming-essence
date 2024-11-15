package main

import "fmt"

func main() {
	// IF文の書き方4種類

	// C言語for形式
	for i := 0; i < 5; i++ {
		fmt.Println("C言語for形式ループ :", i)
	}

	// C言語while形式
	j := 0
	for j < 3 {
		fmt.Println("C言語while形式ループ :", j)
		j++
	}

	// 無限ループ
	// for {
	//	fmt.Println("無限ループ")
	// }

	// map/スライスのイテレート
	s := []string{"Golang", "Java", "Python", "C++"}
	for i, v := range s {
		fmt.Println("No :", i, " Lang :", v)
	}

	// Cと同様にbreak, continueの制御も可能
}
