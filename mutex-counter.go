package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter は並列処理しても安全。
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc は指定されたキーのカウンタをインクリメントする。
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// 一度に 1 つの goroutine だけがマップ c.v. にアクセスできるようにロックする。
	c.v[key]++
	c.mu.Unlock()
}

// Value は指定されたキーのカウンタの現在の値を返す。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// 一度に 1 つの goroutine だけがマップ c.v. にアクセスできるようにロックする。
	defer c.mu.Unlock() // deferを使う理由は、将来if文やpanicが増えた場合に備えてどんな経路でも必ず Unlock される ことを保証するため
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") // 1000個の goroutine が同時に Inc を呼ぶ
	}

	time.Sleep(time.Second) // Sleep は簡易的な待機（本番では sync.WaitGroup を使う）
	fmt.Println(c.Value("somekey"))
}
