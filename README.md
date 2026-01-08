
実務でGoを扱うので、改めて [A Tour of Go](https://go-tour-jp.appspot.com/) を写経しながら、自分用メモとして更新していく。  
初めは頭から書いてたけどやっぱり重要そうなところを先に網羅する。

---

## Basic
### Hello, World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

実行結果：
```
Hello, 世界
```

メモ：
- Go のプログラムは `package main` と `func main()` が必要
- gofmtがgo標準でインストールされているので、ファイルを保存すると自動でコードを整形してくれる
- ビルドは `go build ＜ファイル名＞`
- 実行は `go run ＜ファイル名＞`


---

### Packages

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

実行結果：
```
My favorite number is 6
```

メモ：
- すべての Go ファイルは先頭に package 文を書く
- 同じディレクトリの .go ファイルは原則すべて同じ package 名にしなければならない
- ディレクトリ名とパッケージ名は慣習的に一致させるが、仕様上は一致必須ではない
- エントリポイントは package main 内の func main()


---

### Imports

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g programs.\n", math.Sqrt(7))
}
```

実行結果：
```
Now you have 2.6457513110645907 problems.
```

メモ：
- `import` で複数パッケージをまとめてインポート
- 1行ずつ書くこともできるが、Go のコードスタイルではまとめて書くことが推奨されている



---

### Exported names

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

実行結果：
```
3.141592653589793
```

メモ：
- Goでは、名前の最初の文字で、公開される範囲が決まる
- 大文字で始まる名前 = エクスポート（公開）
  - 外部パッケージから参照可能
  - 例: `math.Pi`, `math.Sqrt`
- 小文字で始まる名前 = 非エクスポート（非公開）
  - 同じパッケージ内でのみ参照可能
  - 例: `math.pi`, `math.hoge`


---

### Functions

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

実行結果：
```
55
```

メモ：
- Goでは引数で、変数名の 後ろ に型名を書く
  - 例：`add(x int)`


---

### Functions continued
```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

実行結果；
```
55
```

メモ：
- 引数の型を省略してまとめて書くことができる `(x, y int)`
- 実務でも2~3個の引数が同じ型ならまとめて可読性上げることもある
- 4つ以上の場合は構造体を定義して、構造体を引数にした方が無難


---

### Multiple results

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

実行結果：
```
world hello
```

メモ：
- Goは複数の戻り値を返すことができる
- 特に以下のように、関数の戻り値を `value, err` で返す書き方はGoのベストプラクティス
    ```
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    ```
- 3つ以上返すのは避けたい。構造体でまとめた方が良い


---

### Named return values

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

実行結果：
```
7 10
```

メモ：
- Goでは戻り値となる変数に名前をつけることができる
- 名前付きの戻り値の変数を使うと、returnで何も書かずに戻せる
- 短い関数以外では使わない方が良い


---

### [未]Variables
### [未]Variables with initializers

---
### Short variable declarations
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

実行結果：
```
1 2 3 true false no!
```

メモ：
- Goでは`var i int = 10`という書き方で変数宣言もできる
- `i := 10`という代入文を使って暗黙的な型宣言もできる
- 関数内では`:=`を積極的に使うべき
- 関数外と、明示的に型指定したい場合に`var`を使う(int64、float64、構造体ゼロ値など)


---

### [未]Basic types

---

### Zero values
```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

実行結果：
```
0 0 false ""
```

メモ：
- 初期値を与えずに変数宣言すると以下のゼロ値が入る
  - 数値型(int,floatなど): `0`
  - bool型: `false`
  - string型: `""`
  - pointer, slice, map, chan, func: `nil`
  - struct: `各フィールドがその型のゼロ値になる`
- ゼロ値を使うべきシーン
  - struct、slice など初期化しなくても普通に動く型
  - カウンタ・一時変数など
- ゼロ値を使うべきではないシーン
  - map のように ゼロ値(nil)だと使えずエラーにつながるもの


---

### [未]Type conversions
### [未]Type inference
### [未]Constants
### [未]Numeric Constants

---

### [未]For
### [未]For continued
### [未]For is Go's "while"
### [未]Forever
### [未]If
### [未]If with a short statement
### [未]If and else
### [未]Exercise: Loops and Functions
### [未]Switch
### [未]Switch evaluation order
### [未]Switch with no condition
### [未]Defer
### [未]Stacking defers
### [未]Congratulations!

---

### Pointers
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // pに、iのアドレスを代入 (var p *int = &i と同じ意味)
	fmt.Println(*p) // *p(iのアドレスの実体)を出力しているので、出力は 42
	*p = 21         // *p(iのアドレスの実体)に21を代入 (=iに上書き)
	fmt.Println(i)  // 上書きされたiを出力しているので、出力は 21

	p = &j         // pに、jのアドレスを代入(上書き)
	*p = *p / j    // *p(jのアドレスの実体)に、*p/j の計算結果を代入 (=jに上書き)
	fmt.Println(j) // 上書きされたjを出力しているので、出力は 1
}
```

実行結果：
```
42
21
1
```

メモ：
| 記法 | 意味 |
| --- | --- |
| `&i` | 変数i のアドレスを取得する |
| `p := &i` | pに変数iのアドレスを代入
| `*p` | pが指しているアドレスの値を取り出す（デリファレンス）
| `*p = 21` | ポインタ経由でアドレスの値を書き換える
|`var p *int` | intを指すポインタ型 を宣言


---

### Structs

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

実行結果：
```
{1 2}
```

メモ：
- struct (構造体)は、フィールドの集まり
- `type 名前 =? 型定義` の構文で型に名前をつける。構造体以外でも使う。
  - 構造体：`type A struct {}`
  - インターフェース：`type B interface {}`
  - 独自型：`type UserID int`
  - 関数：`type Fn func()`
  - スライス/マップ：`type IDs []int`


---

### [未]Struct Fields

---

### Pointers to structs

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

実行結果：
```
{1000000000 2}
```

メモ：
- structのフィールドは、structのポインタを通してアクセスすることもできる
- `p.X`でアクセスできる
   - ちなみに`(*p).X`と書くことでアクセスもできるが長いのであまり使われない

---

### [未]Struct Literals
### [未]Arrays

---

### Slices
```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // 長さ6の固定長配列 primes を宣言して値をセット

	var s []int = primes[1:4]  // 長さ可変のスライス s を宣言して、primesのインデックス1から4未満を参照するようにセット
	fmt.Println(s)
}
```

実行結果：
```
[3 5 7]
```

メモ：
- Goでは、配列`[n]int`は固定長。スライス`[]int`は可変長。



---

### Slices are like references to arrays

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX" // b[0] は names[1] を指しており、元のnames配列を書き換えている
	fmt.Println(a, b) // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}
```

実行結果：
```
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

メモ：
- スライスは配列そのものを持たず、配列の一部を“参照”している
- スライス同士が同じ配列を共有することがある
- どれか1つのスライスを変更すると、配列・他のスライスにも影響する
- コピーしたい場合は`copy()`か`append()`を使う


---

### [未]Slice literals
### [未]Slice defaults

---

### Slice length and capacity
```go
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
```

実行結果：
```
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

メモ：
- スライスは長さ(length)と容量(capacity)の両方を持っている
- 長さ(length)とは
  - `len(s)`で取得できる
  - スライスに「今含まれている要素の数」
  - `for range` で回る範囲
  - `s[i]` でアクセスできる最大範囲
- 容量(capacity)とは
  - `cap(s)`で取得できる
  - スライスの「先頭要素」から、元配列の末尾までの要素数
  - 元配列基準
  - 「この slice をどこまで伸ばせるか」を表す
- 容量を超えて伸ばすと何が起こるか
  - `s = s[:5]` → panic（実行時エラー）


---

### Nil slices

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

実行結果：
```
[] 0 0
nil!
```


メモ：
- スライスのゼロ値は `nil`
- ただし`Println`の表示では空配列と同じく`[]`で表示される
    ```
		var s1 []int            // nil slice
		s2 := []int{}           // 空スライス
		fmt.Println(s1)         // []
		fmt.Println(s2)         // []
		fmt.Println(s1 == nil)  // true
		fmt.Println(s2 == nil)  // false
	```

---

### Creating a slice with make

```go
package main

import "fmt"

func main() {
	a := make([]int, 5) // 長さ5のスライスを作る。すべてゼロ値で初期化。
	printSlice("a", a)  // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5) // 長さ0、容量5のスライスを作る
	printSlice("b", b)     // b len=0 cap=5 []

	c := b[:2]         // 「先頭から2要素分」を 見えるようにした
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]        // 開始位置を2にずらして元配列の index 2〜4 を指すスライスを作った
	printSlice("d", d) // d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

実行結果：
```
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
```

メモ：
- make([]T, len, cap) は 配列を確保し、その配列を参照するスライスを作る
- len は「今使っている長さ」、cap は「将来使える余白」
- cap を指定することで append 時の再確保を防げる
- 高頻度処理・件数が見積もれる場合は make を使う意味がある
- make([]T, 0, n) は「これから詰める」意図を明示できる
- `[:n]` →「先頭から n 個見る」
- `[i:j]` →「i から j-1 まで見る」
- パッと見て分かりづらいのでまとめ直したい


---

### [未]Slices of slices

---

### Appending to a slice

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

実行結果：
```
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```

メモ：
- `append`はスライスに要素を追加する関数
  - 厳密には「新しいスライスを返す関数」
  - `cap`が足りなければ配列を作り直す
- PHPと違いGoにはこの方法しか要素追加する方法はない
  - PHPの場合は`$arr[] = $value`や`array_push()`など複数あるけど
- 戻り値が必要
  - `s = append(s, 1)` このように書く必要がある
  - `append(s, 1)` これだと追加されない

---

### Range

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

実行結果：
```
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

メモ：
- Goの`for + range`は、PHP等の`foreach`とほぼ同じ
- `for + range`は値コピーなので元の値(pow)が書き換えられない
  - PHPもデフォルトは値コピー。オプションで参照にもできるが基本使わないのでほぼ同じと考えてOK


---

### [未]Range continued
### [未]Exercise: Slices

---

### Maps

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

実行結果：
```
{40.68433 -74.39967}
```

メモ：
- Goの連想配列は`map`
- mapのゼロ値は `nil`
- PHPと異なり、`make`しないと要素追加できない(宣言だけでは使えない)
- Goのmapは型安全なのでキーごとに異なる型の値は入れられない(PHPは入れられる)


---

### [未]Map literals
### [未]Map literals continued

---

### Mutating Maps

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42 // 要素の追加
	fmt.Println("The value:", m["Answer"]) // 要素の取得

	m["Answer"] = 48 // 要素の更新
	fmt.Println("The value:", m["Answer"]) // 要素の取得

	delete(m, "Answer") // 要素の削除
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // 存在チェック
	fmt.Println("The value:", v, "Present?", ok)
}
```

実行結果：
```
The value: 48
The value: 42
The value: 0
The value: 0 Present? false
```

メモ：
- mapは 存在しないキーでもゼロ値が返る
- 存在確認が必要な場合は必ず ok を使う
  - valueはゼロ値が入るため存在確認に使えない
- 追加・更新・削除は安全で、存在しないキーに対する操作でもpanicしない


---

### [未]Exercise: Maps

---

### Function values
```go
package main

import (
	"fmt"
	"math"
)

// compute は「2つの float64 を受け取って float64 を返す関数」を受け取り、
// その関数に (3, 4) を渡して結果を返す
func compute(fn func(float64, float64) float64) float64 {
	// fn(3, 4)
	return fn(3, 4)
}

func main() {
	// hypot は直角三角形の斜辺の長さを求める関数
	// 数学的には：√(x² + y²)
	hypot := func(x, y float64) float64 {
		// x*x + y*y  = x² + y²
		// math.Sqrt = √
		return math.Sqrt(x*x + y*y)
	}

	// √(5² + 12²) = √(25 + 144) = √169 = 13
	fmt.Println(hypot(5, 12))

	// compute(hypot)
	// = hypot(3, 4)
	// = √(3² + 4²) = √25 = 5
	fmt.Println(compute(hypot))

	// compute(math.Pow)
	// = math.Pow(3, 4)
	// = 3⁴ = 81
	fmt.Println(compute(math.Pow))
}
```

実行結果：
```
13
5
81
```

メモ：
- Go では 関数も値（変数）として扱える
- 関数は
  - 変数に代入できる
  - 引数として渡せる
  - 戻り値として返せる
- 関数値を使うと 処理の中身を差し替えられる設計ができる
- interface を作らなくても関数だけ渡せる
- DDD / クリーンアーキテクチャで以下によく使われる
  - バリデーション
  - ポリシー差し替え
  - ユースケースの振る舞い切り替え

---

### Function closures

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Println(pos(1))
}
```

実行結果：
```
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
46
```

メモ：
- Go の関数は クロージャ(外側の変数を捕まえて使う関数)になれる  
- このコードでのクロージャは以下(sum を捕まえている無名関数)
  ```go
  func(x int) int {
      sum += x
      return sum
  }
	```
- adder 自体はクロージャではなく、クロージャを生成して返す関数
- sum := 0 は
  - adder() を呼んだ瞬間に1回だけ初期化される
  - クロージャを呼ぶたびに初期化されるわけではない
- adder() を呼ぶたびに別の sum を持つクロージャが作られる
- そのため pos と neg は独立した状態を持つ
	```go
	adder()
	├─ sum = 0
	└─ func(x) { sum += x }  ← クロージャ①（pos）

	adder()
	├─ sum = 0
	└─ func(x) { sum += x }  ← クロージャ②（neg）
	```
- クロージャは「関数 + 捕まえた変数」のセット
- 実務では設定や依存を閉じ込める用途は使って良いが、状態を持ち続ける用途は struct の方が無難

---

### [未]Exercise: Fibonacci closure

---

## Methods and interfaces

### Methods
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

実行結果：
```
5
```

メモ：
- Goにクラスはない
- 型に対してメソッドを定義できる（構造体以外にも独自型にも定義可能）
- このコードでは `(v Vertex)` の部分がレシーバ
- メソッドは `v.Abs()` のように呼べる
- 「データ + 振る舞い」を表現するための仕組み

---

### [スキップ]Methods are functions

---

### Methods continued

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

実行結果；
```
1.4142135623730951
```


メモ：
- Goでは struct に限らず、自分で定義した任意の型 にメソッドを定義できる
（数値型・スライス型・マップ型なども可）
- メソッドは レシーバ付きの関数（型に「振る舞い（できること）」を与える仕組み）
- メソッドを定義できるのは、その型が定義されたパッケージ内だけ。以下には直接メソッドを追加できない
  - 他パッケージの型
  - 組み込み型（int / float64 など）
- 他人の型を拡張したい場合は、ラップ型を作る か embedding（合成）する

---

### Pointer receivers

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

実行結果：
```
50
```

メモ：
- レシーバには2種類ある
  ```go
  func (v Vertex) Method()   // 値レシーバ（構造体のコピーを受け取る）
  func (v *Vertex) Method()  // ポインタレシーバ（構造体そのものを指す）
  ```
- 値レシーバはコピーを操作するので呼び出し元に反映されない
- ポインタレシーバは元の値を変更できる。状態を変更するメソッドでは必須
- 実務ではポインタレシーバで統一されがち
  - 状態を持つ構造体は変更される可能性が高い  
  - 値／ポインタが混在すると理解コストが上がる  
  - そのため すべてポインタレシーバにする設計 が多い  
- パフォーマンス面の利点
  - 大きな構造体を値レシーバにするとコピーコストが発生  
  - ポインタレシーバならコピーなし  
- 元の値を変更したいなら、必ずポインタレシーバ
- Goのメソッドは値渡しが基本のため、構造体の状態を変更するメソッドはポインタレシーバを使う

---

### [スキップ]Pointers and functions
### [未]Methods and pointer indirection
### [未]Methods and pointer indirection (2)
### [未]Choosing a value or pointer receiver

---

### Interfaces

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat は Abser を実装している
	a = &v // *Vertex は Abser を実装している

	// 次の行のように代入すると、v は Vertex 型であり（*Vertex ではないため）
	// Abser インタフェースを実装しておらず、コンパイルエラーになる。
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

実行結果：
```
5
```

メモ：
- interface は メソッドのシグネチャ(名前・引数・戻り値)の集合
- interface はGoでは明示的に implements しない（暗黙実装）
- interface 変数には要求メソッドをすべて持つ型の値だけ代入できる
- interface の実装可否は`method set`（どのメソッドを **静的に** 持っているか）で判定される
  - 型T と 型*T では method set が違う。
  - そのため、上のコードでは
    - `a = f` （MyFloat 型） → コンパイル成功（`Abs()`メソッドを持っているため）
    - `a = &v` （*Vertex 型） → コンパイル成功（`Abs()`メソッドを持っているため）
    - `a = v` （Vertex 型） → コンパイルエラー（`Abs()`メソッドを持っていないため）
- 実務では、ポインタレシーバを使うことで method set 問題を回避できるので、`var i IFace = &Struct{}` が基本形

---

### [スキップ]Interfaces are implemented implicitly

---

### Interface values

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"} // 内部的には (value = &T{"Hello"}, type = *T)
	describe(i)
	i.M()

	i = F(math.Pi) // 内部的には (value = F(3.1415...), type = F)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

実行結果：
```
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

メモ：
- interface の値は (値, 具体型) のペアで構成される
- interface は 実行時にも具体型の情報を保持している
- interface 変数には 異なる具体型の値を代入できる
- interface のメソッド呼び出しは保持している具体型のメソッドが動的に呼ばれる
- **同じ interface 変数でも、代入する値によって中身の具体型と振る舞いが変わる**
  - この仕組みが 多態性（ポリモーフィズム）を実現している

---

### Interface values with nil underlying values

```go
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
	// i.M()    // 具体型が入っていないため、panicになる

	var t *T // t は nil だが、*T という型情報が interface に入る
	i = t
	describe(i) // (<nil>, *main.T) // 具体値がnil、型が*main.T
	i.M()       // <nil> // 型があるため、nil レシーバーとしてメソッドを呼び出せる。
	            // if t == nil の条件に引っかかるため、<nil> が出力される

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T) // 具体値が&T{"hello"}、型が*main.T
	i.M()       // hello // if t == nil の条件に引っ掛からずfmt.Println(t.S)が実行される
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

実行結果：
```
(<nil>, <nil>)
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```

メモ：
- interface は (具体型, 具体値) の2つの情報を内部に持っている
- 型も値も nil のときだけ if i == nil が true
- nil ポインタを interface に代入すると
  - 具体型は入る
  - 具体値は nil
  - interface 自体は非 nil になる
- そのため if i == nil が false になるケースがある
- 具体値が nil でも メソッド呼び出しは可能（nil レシーバー）
- interface（特に error）を返す関数は成功時に必ず return nil しないと事故りがち

---

### [スキップ]Nil interface values
### [スキップ]The empty interface
### [スキップ]Type assertions
### [スキップ]Type switches

---

### [未]Stringers

```go

```

実行結果：
```

```

メモ：
- 


---


### [未]Exercise: Stringers
### [未]Errors
### [未]Exercise: Errors
### [未]Readers
### [未]Exercise: Readers
### [未]Exercise: rot1 3Reader
### [未]Images
### [未]Exercise: Images

---

## Generics
- 日本語版にはまだ無いページっぽい（英語版:https://go.dev/tour/generics/1）

### Type parameters

```go
package main

import "fmt"

// Index は s 内の x のインデックスを返す。見つからない場合は -1 を返す
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		//  v と x はcomparable型 T なので、ここでは == を使用できる
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
```

実行結果：
```
2
-1
```

メモ：
- ジェネリクスとは、**型違いでコピペしていたコードを、1つにまとめるための仕組み**
- Goのジェネリクス（Type Parameters）とは
  - ジェネリクス = 型を引数として受け取る仕組み
  - 実行時ではなく コンパイル時に型が確定する
  - 型パラメータは関数名の直後に書く
    ```go
    func Foo[T any](...)
    ```
- 型パラメータと制約
  - T は 型の変数
  - any = 制約なし（= interface{} の別名）
  - 制約（constraint）を書くことで使える演算・操作をコンパイラに保証させる
- comparable
  - comparable は `==` `!=` が使える型のみ許可
  - OK: int, string, bool, pointer, comparableなstruct
  - NG: slice, map, func
  - 制約がないと比較演算子は使えない
- 上のコードの例
  - `Index`関数は比較可能な型のスライスから値を探す関数
  - s と x は 必ず同じ型 T
  - 型ごとに関数を量産しなくてよい(ジェネリクスの恩恵)
- ジェネリクスと インターフェース の違い
  - ジェネリクス: 型の抽象化（何の型か）
  - インターフェース: 振る舞いの抽象化（何ができるか）
  - ジェネリクスは静的（コンパイル時）
  - インターフェース は動的（実行時）

---

## Generic types

```go
package main

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
```

実行結果：
```
出力なし
```

メモ：
- ジェネリクスは 関数だけでなく構造体（struct）にも使える
- T は struct 全体で共有される型パラメータ
- フィールドにも T をそのまま使える
  - 値保持用：val T
  - 再帰参照：*List[T]
- 上のコードでは型安全な再帰構造を作っている
- 実務では自作することは少ない

---

## Concurrency
### [未]Goroutines
### [未]Channels
### [未]Buffered Channels
### [未]Range and Close
### [未]Select
### [未]Default Selection
### [未]Exercise: Equivalent Binary Trees
### [未]Exercise: Equivalent Binary Trees sync.Mutex
### [未]Exercise: Web Crawler

---
