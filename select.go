package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1 // フィボナッチ数列の初期値
	for {        // 無限ループ。 select内で止める前提
		select {
		case c <- x: // 受信側がいて、c に送信できるなら x を送る
			x, y = y, x+y
		case <-quit: // quit に何か送られてきたら"quit" を表示して関数を終了
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // c から 10 回受信して表示
		}
		quit <- 0 // 0 に意味はなく「送られた」という事実が重要
	}()
	fibonacci(c, quit)
}
