package main

import (
	"fmt"
	"log"
)

// 条件分岐
func main() {

	// if文はかっこを必要としない
	x := 1
	y := 2
	if x == 2 && y == 3 {
		//do something
	}

	// セミコロンを用いてifの範囲内で使用する変数宣言も可能
	// この書き方ではセミコロン以降がifの判定対象
	if user, err := userName(); err == nil {
		fmt.Println(user)
	} else {
		// ifのなかで宣言した変数やerrはelseでも利用可能
		log.Println(err)
	}

	// switchはbreakを書かなくても、1つのケースに一致した後に処理を抜ける
	i := 11
	switch i {
	case 1:
		fmt.Println("one")

	case 11:
		fmt.Println("eleven")
		// 次のcaseに降下したい場合は明示的にfallthroughを使う
		// この場合、i = 100に合致しなくとも、case 100が実行される。
		fallthrough

	case 100:
		fmt.Println("hundred")
	}

	// caseには固定値ではなく式を書くことができる
	z := 10
	hit := 100
	switch {
	case z > hit:
		fmt.Println("bigger than", hit)
	case z < hit:
		fmt.Println("less than", hit)
	case z == hit:
		fmt.Println("equal to", hit)
	}

}

// サンプル用関数
func userName() (user string, err error) {
	return "satoshi", nil
}
