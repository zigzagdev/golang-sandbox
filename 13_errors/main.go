package main

import (
	"errors"
	"fmt"
)

// ----------------------------------------------------------------
// A Tour of Go: Errors
// https://go.dev/tour/methods/19
// ----------------------------------------------------------------

// Goのエラーは error インターフェースを実装した値
// PHP: 例外（throw/catch）と違い、戻り値として扱う

// error インターフェース:
// type error interface {
//     Error() string
// }

// カスタムエラー型
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// 関数はエラーを最後の戻り値として返す（Goのイディオム）
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// カスタムエラーを返す
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "must be <= 150"}
	}
	return nil
}

// fmt.Errorf でエラーをラップする（%w でアンラップ可能）
func fetchUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("fetchUser: invalid id %d: %w", id, errors.New("id must be positive"))
	}
	return nil
}

func main() {
	// 基本的なエラーハンドリング
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result) // 5
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println(err) // division by zero
	}

	// カスタムエラー
	if err := validateAge(-1); err != nil {
		fmt.Println(err) // validation error: age - must be non-negative
	}

	// errors.As でカスタムエラー型を取り出す
	err = validateAge(-1)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("Field:", ve.Field)   // Field: age
		fmt.Println("Msg:", ve.Message)   // Msg: must be non-negative
	}

	// errors.Is でエラーチェーンを確認する
	sentinel := errors.New("sentinel error")
	wrapped := fmt.Errorf("wrapped: %w", sentinel)
	fmt.Println(errors.Is(wrapped, sentinel)) // true

	// ラップされたエラー
	err = fetchUser(-1)
	fmt.Println(err) // fetchUser: invalid id -1: id must be positive
}
