package main

import "fmt"

// ----------------------------------------------------------------
// A Tour of Go: Channels
// https://go.dev/tour/concurrency/2
// ----------------------------------------------------------------

// channel：goroutine 間でデータを送受信するパイプ
// <- 演算子で送受信する
// make(chan 型) / make(chan 型, バッファサイズ)

func sum(s []int, ch chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	ch <- total // チャネルに送信
}

// ----------------------------------------------------------------
// Range and Close（A Tour of Go: Range and Close）
// ----------------------------------------------------------------

func generate(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch) // 送信完了を通知（受信側で range が終了する）
}

// ----------------------------------------------------------------
// Select（A Tour of Go: Select）
// 複数のチャネル操作を待ち、準備できた方を実行する
// switch の channel 版
// ----------------------------------------------------------------

func fibonacci(ch, quit chan int) {
	a, b := 0, 1
	for {
		select {
		case ch <- a: // ch に送れるなら送る
			a, b = b, a+b
		case <-quit: // quit から受信したら終了
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 基本的なチャネル（バッファなし：送受信が揃うまでブロック）
	ch := make(chan int)
	s := []int{7, 2, 8, -9, 4, 0}

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x, y := <-ch, <-ch // 両方の goroutine から受信
	fmt.Println(x, y, x+y) // -5 17 12

	// バッファ付きチャネル（バッファが満杯になるまでブロックしない）
	bch := make(chan int, 3)
	bch <- 1
	bch <- 2
	bch <- 3
	fmt.Println(<-bch) // 1
	fmt.Println(<-bch) // 2

	// range でチャネルを受信（close されるまで繰り返す）
	numCh := make(chan int, 5)
	go generate(5, numCh)
	for n := range numCh {
		fmt.Print(n, " ") // 0 1 2 3 4
	}
	fmt.Println()

	// select
	fibCh := make(chan int, 10)
	quit := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			fmt.Print(<-fibCh, " ") // フィボナッチ数列を受信
		}
		fmt.Println()
		quit <- 0 // 終了を通知
	}()

	fibonacci(fibCh, quit) // 0 1 1 2 3 5 8 13 \nquit
}
