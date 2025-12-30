
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
### [未]Function values
### [未]Function closures
### [未]Exercise: Fibonacci closure

---

## Methods and interfaces
### [未]Methods and interfaces
### [未]Methods
### [未]Methods are functions
### [未]Methods continued
### [未]Pointer receivers
### [未]Pointers and functions
### [未]Methods and pointer indirection
### [未]Methods and pointer indirection (2)
### [未]Choosing a value or pointer receiver
### [未]Interfaces
### [未]Interfaces are implemented implicitly
### [未]Interface values
### [未]Interface values with nil underlying values
### [未]Nil interface values
### [未]The empty interface
### [未]Type assertions
### [未]Type switches
### [未]Stringers
### [未]Exercise: Stringers
### [未]Errors
### [未]Exercise: Errors
### [未]Readers
### [未]Exercise: Readers
### [未]Exercise: rot1 3Reader
### [未]Images
### [未]Exercise: Images

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
