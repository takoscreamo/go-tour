package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil { // interface の中の具体値が nil のケース
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I     // 型も値も入っていない → interface は nil
	describe(i) // (<nil>, <nil>) // 具体値がnil、型もnil

	var t *T // t は nil だが、*T という型情報が interface に入る
	i = t
	describe(i) // (<nil>, *main.T) // 具体値がnil、型が*main.T
	i.M()       // <nil> // if t == nil の条件に引っかかるため、<nil> が出力される

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T) // 具体値が&T{"hello"}、型が*main.T
	i.M()       // hello // if t == nil の条件に引っ掛からずfmt.Println(t.S)が実行される
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
