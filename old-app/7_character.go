package main

import "fmt"

func main() {
	// 文字列

	// 文字列の結合
	s := "Hello "
	name := "sakats"
	s += name
	fmt.Println(s)

	// 文字列はバイト列のため添え字でアクセス可
	fmt.Printf("%c", s[0])

	// stringはイミュータブルなので一部を書き換える場合はバイトに置換
	b := []byte(s)
	b[0] = 'h'    // Helloをhelloに置換
	s = string(b) // イミュータブルなので再代入
	fmt.Println(s)

	// Goにはruneという型があり、Unicodeのコードポイントに変換可
	st := "こんにちわ世界"
	rs := []rune(st)
	rs[4] = 'は'     //「わ」を「は」に変換
	st = string(rs) // イミュータブルなので再代入
	fmt.Println(st)

	//固定長の配列は[:]という指定でスライスに変換できる
	var bi [4]byte
	bi2 := bi[:] // bi2はスライス
	fmt.Println(bi2)

	// 複数行テキスト
	var content = `これは
	複数行からなる
	コンテンツです
	バックスラッシュをエスケープしないで出力できる
	\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	`
	fmt.Println(content)

}
