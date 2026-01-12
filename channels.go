package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 計算結果をチャネルへ送信（戻り値なし）
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 前半: [7, 2, 8]
	go sum(s[len(s)/2:], c) // 後半: [-9, 4, 0]
	x, y := <-c, <-c        // 先に終わった goroutine の値が x に入る(順不同)

	fmt.Println(x, y, x+y)
}
