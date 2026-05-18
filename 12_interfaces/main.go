package main

import (
	"fmt"
	"math"
)

// ----------------------------------------------------------------
// A Tour of Go: Interfaces
// https://go.dev/tour/methods/9
// ----------------------------------------------------------------

// インターフェース：メソッドのシグネチャの集合
// Go のインターフェースは「暗黙的」に実装される
// PHP: implements キーワードが必要 → Go: 宣言不要

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func printShape(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ----------------------------------------------------------------
// Stringer インターフェース（A Tour of Go: Stringers）
// fmt.Println はこのインターフェースを自動で呼ぶ
// PHP: __toString() に相当
// ----------------------------------------------------------------

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// ----------------------------------------------------------------
// 空インターフェース（A Tour of Go: The empty interface）
// any / interface{} はすべての型を受け入れる
// ----------------------------------------------------------------

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ----------------------------------------------------------------
// 型アサーション（A Tour of Go: Type assertions）
// interface の値から具体的な型の値を取り出す
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// 型スイッチ（A Tour of Go: Type switches）
// ----------------------------------------------------------------

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}

func main() {
	// 暗黙的なインターフェース実装
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 3},
	}
	for _, s := range shapes {
		printShape(s)
	}

	// Stringer
	p := Point{3, 4}
	fmt.Println(p) // (3, 4)

	// 空インターフェース
	describe(42)
	describe("hello")
	describe(true)

	// 型アサーション
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	n, ok := i.(int)
	fmt.Println(n, ok) // 0 false

	// 型スイッチ
	typeSwitch(42)
	typeSwitch("go")
	typeSwitch(true)
	typeSwitch(3.14)
}
