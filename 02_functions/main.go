package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Functions
// https://go.dev/tour/basics/4
// ----------------------------------------------------------------

// 基本的な関数：型は変数名の後に書く（PHPと逆）
// PHP: function add(int $x, int $y): int
func add(x int, y int) int {
	return x + y
}

// 同じ型が連続するとき、最後の1つだけ書ける
func multiply(x, y int) int {
	return x * y
}

// 複数の戻り値（A Tour of Go: Multiple Results）
// PHPでは配列で代替するが、Goはネイティブでサポート
func swap(a, b string) (string, string) {
	return b, a
}

// Named return values（A Tour of Go: Named Return Values）
// 戻り値に名前をつけると、return だけで返せる（naked return）
// ※ 短い関数向け。長い関数では可読性が落ちる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 可変長引数（variadic）
// ...型 で任意個の引数を受け取れる
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(add(3, 4))       // 7
	fmt.Println(multiply(3, 4))  // 12

	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	x, y := split(17)
	fmt.Println(x, y) // 7 10

	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	// スライスを展開して渡すこともできる
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...)) // 10
}
