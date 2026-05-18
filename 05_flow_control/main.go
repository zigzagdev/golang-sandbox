package main

import (
	"fmt"
	"math"
)

// ----------------------------------------------------------------
// A Tour of Go: Flow Control Statements
// https://go.dev/tour/flowcontrol/1
// ----------------------------------------------------------------

func forExamples() {
	// Goのforは ( ) が不要（PHPと違い {} は必須）
	// PHP: for ($i = 0; $i < 3; $i++) {}
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// for は while の代わりにもなる（条件だけ書く）
	// PHP: while ($n < 100) {}
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n) // 128

	// 無限ループ
	// for { ... } // break で抜ける
}

func ifExamples() {
	// if の短文（short statement）：条件の前に文を書ける
	// スコープは if/else ブロック内に限定される
	if x := math.Sqrt(16); x < 5 {
		fmt.Println("small:", x) // small: 4
	} else {
		fmt.Println("big:", x)
	}
	// ここでは x は使えない

	// PHPの else if は Go では else if（1語ではなく2語）
}

func switchExamples() {
	// Goの switch は条件が true になった case だけ実行する
	// PHPと違い break は不要（fallthrough を明示すると次のcaseも実行）
	x := 3
	switch x {
	case 1:
		fmt.Println("one")
	case 2, 3: // カンマ区切りで複数条件
		fmt.Println("two or three")
	default:
		fmt.Println("other")
	}

	// 条件なし switch（if-else chain の代替として読みやすい）
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("morning")
	case hour < 18:
		fmt.Println("afternoon") // afternoon
	default:
		fmt.Println("evening")
	}
}

func deferExamples() {
	// defer：関数リターン時に実行を遅延させる
	// ファイルのクローズ、ロック解放など後処理に使う
	// PHP: finally ブロックに近い

	defer fmt.Println("world") // 後で実行

	fmt.Println("hello") // 先に実行

	// defer はスタック（LIFO）で積まれる
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
	// 出力は後で: 2 1 0
}

func main() {
	fmt.Println("--- for ---")
	forExamples()

	fmt.Println("--- if ---")
	ifExamples()

	fmt.Println("--- switch ---")
	switchExamples()

	fmt.Println("--- defer ---")
	deferExamples()
	// deferExamples の defer が main 終了後に実行される
}
