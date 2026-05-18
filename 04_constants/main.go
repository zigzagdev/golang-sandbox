package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Constants
// https://go.dev/tour/basics/15
// ----------------------------------------------------------------

// const で定数を定義する（:= は使えない）
const Pi = 3.14159

// 複数まとめて定義
const (
	StatusOK    = 200
	StatusNotFound = 404
)

// ----------------------------------------------------------------
// iota：連番の定数を自動生成する
// A Tour of Go では Numeric Constants として登場
// ----------------------------------------------------------------

type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// iota を使った bit flags（実用的なパターン）
type Permission uint

const (
	Read    Permission = 1 << iota // 1 (001)
	Write                          // 2 (010)
	Execute                        // 4 (100)
)

// ----------------------------------------------------------------
// Untyped constants（型なし定数）
// 精度が高く、文脈に応じて型が決まる
// ----------------------------------------------------------------

const Big = 1 << 62
const Small = Big >> 61 // 2

func needInt(x int) int         { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(Pi)          // 3.14159
	fmt.Println(StatusOK)    // 200
	fmt.Println(StatusNotFound) // 404

	fmt.Println(North, East, South, West) // North East South West

	perm := Read | Write
	fmt.Println(perm)            // 3
	fmt.Println(perm&Read != 0)  // true（Read権限あり）
	fmt.Println(perm&Execute != 0) // false（Execute権限なし）

	// 型なし定数は文脈によって int にも float64 にもなる
	fmt.Println(needInt(Small))   // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 4.611686018427388e+17
}
