package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Maps
// https://go.dev/tour/moretypes/19
// ----------------------------------------------------------------

// マップはキーと値のペア。PHPの連想配列に相当
// PHP: $m = ["key" => "value"];
// Go:  m := map[string]string{"key": "value"}

type Vertex struct {
	Lat, Long float64
}

func main() {
	// make で初期化（nil マップへの書き込みは panic になる）
	m := make(map[string]int)

	// 追加・更新
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // map[one:1 two:2]

	// 読み取り
	fmt.Println(m["one"]) // 1

	// 削除
	delete(m, "two")
	fmt.Println(m) // map[one:1]

	// カンマOKイディオム：キーが存在するか確認
	// PHP: array_key_exists($key, $arr)
	val, ok := m["one"]
	fmt.Println(val, ok) // 1 true

	val, ok = m["missing"]
	fmt.Println(val, ok) // 0 false（ゼロ値が返る）

	// マップリテラル
	locations := map[string]Vertex{
		"Tokyo":    {35.6895, 139.6917},
		"New York": {40.7128, -74.0060},
	}
	fmt.Println(locations["Tokyo"]) // {35.6895 139.6917}

	// ゼロ値は nil（make か リテラルで初期化が必要）
	var nilMap map[string]int
	fmt.Println(nilMap == nil)  // true
	fmt.Println(nilMap["key"]) // 0（読み取りはOK）
	// nilMap["key"] = 1 // panic: assignment to entry in nil map

	// ネストしたマップ
	nested := map[string]map[string]int{
		"a": {"x": 1, "y": 2},
		"b": {"x": 3, "y": 4},
	}
	fmt.Println(nested["a"]["x"]) // 1
}
