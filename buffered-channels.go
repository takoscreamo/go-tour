package main

import "fmt"

func main() {
	// ×悪い例（`ch <- 3`で送信がブロックし、main goroutine が進めなくなってデッドロックする）
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// ◯良い例（バッファを空ける）
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch) // 1つ空く
	// ch <- 3           // ここで送信できる
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// ◎より良い例（goroutineを使って送信と受信を分離）
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
