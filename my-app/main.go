package main

import "fmt"

func main() {
	// 配列の宣言・代入（固定長である）
	var a [4]int
	a[0] = 1
	fmt.Println(a) // 代入を行わない要素はintの初期値0になる

	// スライスの宣言１（可変長である）
	var b []int
	fmt.Println(b) // 初期値はnil

	// スライスの宣言２（makeで行う例）
	c := make([]int, 3)
	c[0] = 1
	c[1] = 2
	c[2] = 3

	// スライスの宣言３（初期化も行う例）
	d := []int{1, 2, 3}
	fmt.Println(d)

	// 多次元の配列
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr1)

	// 多次元のスライス
	arr2 := [][]int{
		{10, 20, 30},
		{40, 50, 60},
	}
	fmt.Println(arr2)

	// スライスの拡張
	s := []int{1, 2, 3}
	fmt.Println(s)
	s = append(s, 4) // appendを用いてスライスを再作成・格納
	fmt.Println(s)
	s = append(s, 5, 6) // 複数値の代入も可能
	fmt.Println(s)
	fmt.Println("aの長さは", len(s))

	// 最大容量を表す「cap」
	// 配列およびスライスにはlenとは異なり、capという容量を表す値がある。
	// 配列の場合はlen=cap, スライスの場合はlen<capと予約分のスペースがある。
	// appendするたびにメモリ確保を行うと実行効率が下がるため、サイズが把握できているなら指定しておく。
	// makeを使うと、lenとcapを指定してスライスを宣言できる。
	arr3 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr3 = append(arr3, i)
	}

	// スライスは添え字で範囲を指定して部分参照できる
	fmt.Println(arr3[2:5])

	// スライスから要素を削除する方法１
	// 新しく同じ型のスライスを用意して詰めなおす
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x2 := make([]int, 0, len(x1)/2)
	for i := 0; i < len(x1); i++ {
		if i%2 == 0 {
			x2 = append(x2, x1[i]) // 要素番号が偶数のモノのみ新しいスライスにappend
		}
	}
	x1 = x2
	fmt.Println(x1)

	// スライスから要素を削除する方法２
	// 部分参照とcopyを使う(要素3を削除)
	n := 3
	x1 = x1[:n+copy(x1[n:], x1[n+1:])]
	fmt.Println(x1)

}
