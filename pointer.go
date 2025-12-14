package main

import "fmt"

func main() {
	i, j := 42, 2701

	var p *int = &i // pに、iのアドレスを代入
	fmt.Println(*p) // *p(iのアドレスの実体)を出力しているので、出力は 42
	*p = 21         // *p(iのアドレスの実体)に21を代入 (=iに上書き)
	fmt.Println(i)  // 上書きされたiを出力しているので、出力は 21

	p = &j         // pに、jのアドレスを代入(上書き)
	*p = *p / j    // *p(jのアドレスの実体)に、*p/j の計算結果を代入 (=jに上書き)
	fmt.Println(j) // 上書きされたjを出力しているので、出力は 1
}
