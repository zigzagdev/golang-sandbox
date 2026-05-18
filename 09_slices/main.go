package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Slices
// https://go.dev/tour/moretypes/7
// ----------------------------------------------------------------

// スライスは配列への参照。長さ可変。Goで最もよく使うコレクション型
// PHP の配列に最も近い

func main() {
	// スライスリテラル（内部で配列を作り、それへの参照を返す）
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s) // [1 2 3 4 5]

	// スライス操作 [low:high]（high は含まない）
	fmt.Println(s[1:3]) // [2 3]
	fmt.Println(s[:2])  // [1 2]
	fmt.Println(s[3:])  // [4 5]

	// len と cap
	// len: 現在の要素数
	// cap: 内部配列の先頭からの長さ（再アロケーションなしで拡張できる上限）
	t := s[1:3]
	fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), t) // len=2 cap=4 [2 3]

	// ゼロ値は nil（len=0, cap=0）
	var nilSlice []int
	fmt.Println(nilSlice == nil) // true

	// make で長さ・容量を指定して作る
	a := make([]int, 3)      // len=3, cap=3
	b := make([]int, 3, 10)  // len=3, cap=10
	fmt.Printf("len=%d cap=%d\n", len(a), cap(a)) // 3 3
	fmt.Printf("len=%d cap=%d\n", len(b), cap(b)) // 3 10

	// append：要素追加。cap が足りなければ自動で再アロケーション
	// PHP: array_push や $arr[] =
	c := []int{1, 2}
	c = append(c, 3, 4)
	fmt.Println(c) // [1 2 3 4]

	// range：スライスをループ（index, value のペアを返す）
	// PHP: foreach ($arr as $i => $v)
	words := []string{"go", "is", "fun"}
	for i, v := range words {
		fmt.Printf("%d: %s\n", i, v)
	}

	// _ でインデックスを捨てる
	total := 0
	for _, n := range []int{1, 2, 3, 4, 5} {
		total += n
	}
	fmt.Println(total) // 15

	// copy：スライスのコピー（参照ではなく値をコピー）
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 99
	fmt.Println(src) // [1 2 3]（影響なし）
	fmt.Println(dst) // [99 2 3]
}
