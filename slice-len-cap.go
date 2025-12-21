package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13} // length=6、capacity=6の配列のスライス
	printSlice(s)                  // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]     // lenを0にする。capは6のまま。
	printSlice(s) // len=0 cap=6 []

	s = s[:4]     // lenを4伸ばす。capは6のまま。
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]     // sliceの先頭ポインタをindex=2にずらす。capが2減って4になる。
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
