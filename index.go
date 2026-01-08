package main

import "fmt"

// Index は s 内の x のインデックスを返します。見つからない場合は -1 を返す
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		//  v と x は比較制約を持つ型 T なので、ここでは == を使用できる
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// インデックスは int のスライスに対して機能する
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// インデックスは文字列のスライスに対しても機能する
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
