package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Generics
// https://go.dev/tour/generics/1
// ----------------------------------------------------------------

// Go 1.18 で追加。型パラメータで再利用可能な関数・型を書ける
// PHP: テンプレートなし（型は動的）

// ----------------------------------------------------------------
// 型パラメータを持つ関数（A Tour of Go: Type Parameters）
// ----------------------------------------------------------------

// comparable: == / != が使える型の制約
func indexOf[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

// 型制約でメソッドを要求する
type Number interface {
	int | int64 | float64
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// ----------------------------------------------------------------
// ジェネリックな型（A Tour of Go: Generic Types）
// ----------------------------------------------------------------

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Len() int { return len(s.items) }

// Map / Filter / Reduce（関数型スタイルのユーティリティ）
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func main() {
	// 型推論で型パラメータを省略できる
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(indexOf(ints, 3))          // 2
	fmt.Println(indexOf(ints, 99))         // -1

	strs := []string{"go", "is", "fun"}
	fmt.Println(indexOf(strs, "is"))       // 1

	// Number 制約
	fmt.Println(sum([]int{1, 2, 3}))       // 6
	fmt.Println(sum([]float64{1.1, 2.2})) // 3.3

	// ジェネリックなスタック
	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Len())  // 3
	v, ok := intStack.Pop()
	fmt.Println(v, ok)          // 3 true

	// 文字列スタックも同じ型で使える
	var strStack Stack[string]
	strStack.Push("hello")
	top, _ := strStack.Pop()
	fmt.Println(top) // hello

	// Map
	doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	fmt.Println(doubled) // [2 4 6]

	lengths := Map([]string{"go", "generics"}, func(s string) int { return len(s) })
	fmt.Println(lengths) // [2 8]
}
