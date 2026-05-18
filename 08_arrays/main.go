package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Arrays
// https://go.dev/tour/moretypes/6
// ----------------------------------------------------------------

// Goの配列は長さが型の一部（固定長）
// PHP: $arr = [1, 2, 3]; // 可変長
// Go:  var arr [3]int    // 長さ3で固定

func main() {
	// var 宣言（ゼロ値で初期化される）
	var a [3]int
	fmt.Println(a) // [0 0 0]

	// リテラルで初期化
	b := [3]string{"foo", "bar", "baz"}
	fmt.Println(b[0], b[1], b[2]) // foo bar baz

	// 長さは len() で取得
	fmt.Println(len(b)) // 3

	// 2次元配列
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix[1][2]) // 6

	// 配列はコピーされる（スライスと異なる）
	c := b
	c[0] = "changed"
	fmt.Println(b[0]) // foo（元は変わらない）
	fmt.Println(c[0]) // changed

	// ... で要素数を自動カウント
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println(len(primes)) // 5
}
