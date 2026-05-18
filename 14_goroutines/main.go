package main

import (
	"fmt"
	"sync"
	"time"
)

// ----------------------------------------------------------------
// A Tour of Go: Goroutines
// https://go.dev/tour/concurrency/1
// ----------------------------------------------------------------

// goroutine：Goランタイムが管理する軽量スレッド
// go キーワードで非同期に関数を実行する
// PHP: 並行処理はなし（pcntl_fork や ReactPHP が必要）

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done() // goroutine 終了時に WaitGroup のカウンタを -1
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

// ----------------------------------------------------------------
// sync.Mutex：共有データの排他制御
// A Tour of Go: sync.Mutex
// https://go.dev/tour/concurrency/9
// ----------------------------------------------------------------

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// WaitGroup で goroutine の完了を待つ
	var wg sync.WaitGroup

	wg.Add(2)
	go say("hello", &wg)
	go say("world", &wg)
	wg.Wait() // 両方の goroutine が Done() するまでブロック

	fmt.Println("--- done ---")

	// sync.Mutex で並行アクセスを安全にする
	c := SafeCounter{v: make(map[string]int)}
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			c.Inc("key")
		}()
	}
	wg2.Wait()
	fmt.Println(c.Value("key")) // 100（レースコンディションなし）
}
