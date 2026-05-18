package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Function Values & Closures
// https://go.dev/tour/moretypes/24
// ----------------------------------------------------------------

// 関数は第一級の値（変数に代入・引数に渡せる）
// PHP: Closure / callable に相当

func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

// クロージャ：外のスコープの変数を「束縛」した関数
// 呼び出すたびに束縛した変数が更新される

func makeCounter() func() int {
	count := 0
	return func() int {
		count++ // 外の count を参照・更新
		return count
	}
}

// アダー（加算器）を生成するファクトリ
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// フィボナッチ数列をクロージャで生成
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		v := a
		a, b = b, a+b
		return v
	}
}

func main() {
	// 関数を変数に代入
	double := func(x int) int { return x * 2 }
	fmt.Println(double(5))            // 10
	fmt.Println(applyTwice(double, 3)) // 12

	// クロージャのカウンター
	// PHP: static 変数に相当するが、インスタンスごとに独立
	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// 別のカウンターは独立した状態を持つ
	counter2 := makeCounter()
	fmt.Println(counter2()) // 1（別の状態）

	// アダー
	add5 := makeAdder(5)
	add10 := makeAdder(10)
	fmt.Println(add5(3))  // 8
	fmt.Println(add10(3)) // 13

	// フィボナッチ
	fib := fibonacci()
	for i := 0; i < 8; i++ {
		fmt.Print(fib(), " ") // 0 1 1 2 3 5 8 13
	}
	fmt.Println()
}
