package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Pointers
// https://go.dev/tour/moretypes/1
// ----------------------------------------------------------------

// ポインタ：変数のメモリアドレスを保持する値
// PHP/Rubyにはポインタがない。Goでは値のコピーか参照かを明示する

func increment(n *int) {
	// * で参照先の値にアクセス（デリファレンス）
	*n++
}

func noEffect(n int) {
	// 値のコピーなので呼び出し元には影響しない
	n++
}

type Point struct {
	X, Y int
}

func main() {
	i := 42

	// & で変数のアドレスを取得する
	p := &i
	fmt.Println(*p) // 42（ポインタを通じて値を読む）

	*p = 100        // ポインタを通じて値を書き換える
	fmt.Println(i)  // 100

	// ゼロ値は nil
	var ptr *int
	fmt.Println(ptr) // <nil>

	// 値渡し vs ポインタ渡し
	x := 10
	noEffect(x)
	fmt.Println(x) // 10（変わらない）

	increment(&x)
	fmt.Println(x) // 11（変わる）

	// struct のポインタ
	// (*v).X と書くのが本来だが、v.X と省略できる（Go の自動デリファレンス）
	v := &Point{1, 2}
	v.X = 99
	fmt.Println(v) // &{99 2}

	// new() でゼロ値ポインタを作る
	q := new(int)
	fmt.Println(*q) // 0
	*q = 5
	fmt.Println(*q) // 5
}
