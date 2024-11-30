package main

import (
	"fmt"
	"sort"
)

func main() {
	// map
	// 順を持たないキーと値のペア。キーは文字列以外の型も利用可能。

	// 下記の宣言はstring型のキー、int型の値を持つmap
	// 初期値はnilのため、mapを作成するためのmakeを使う
	var m map[string]int
	m = make(map[string]int)

	// スライスと同様にcapも指定可能(メモリ制限や高速化になる)
	// m := make(map[string]int, 10)

	// 値の格納は下記の通り
	m["Satoshi"] = 26
	m["Ken"] = 25
	m["T"] = 32

	// 初期化はリテラルを使って初期値の代入をすることでも可
	m2 := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}
	fmt.Println(m2)

	// mapから要素を削除する
	delete(m2, "Bob")
	fmt.Println(m2)

	// mapのキーと値を列挙する
	// mapは順序を持たないため、このfor-rangeは毎回異なる結果になる
	for key, value := range m {
		fmt.Printf("key: %v, value %v\n", key, value)
	}

	// キーでソートしたい場合、先にキーを取り出してソートする必要がある
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // キーを取り出した配列をソート
	for _, k := range keys {
		fmt.Printf("key: %v, value %v\n", k, m[k]) // keyを使って元のmapにアクセス
	}

	// mapは存在しないキーを指定すると、値の型のゼロ値が帰る。
	// intだった場合0だが、初期値0か、検索エラー0かは判定不可
	fmt.Println("key:Akatsukaの検索結果 ", m["Akatsuka"])

	// キーの存在チェックをする場合には、boolの戻り値を受け取って判定する
	val, ok := m["Satoshi"]
	if ok {
		fmt.Printf("key: Satoshi は存在します。値は%vです。", val)
	}

}
