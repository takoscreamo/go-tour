実務でGoを扱うので、[A Tour of Go](https://go-tour-jp.appspot.com/) を写経しながらポイントをまとめました。  
演習ページや重要度の低いページはスキップしてます。  

---
**目次**
- [Basic](#basic)
	- [Hello, World](#hello-world)
	- [Packages](#packages)
	- [Imports](#imports)
	- [Exported names](#exported-names)
	- [Functions](#functions)
	- [Functions continued](#functions-continued)
	- [Multiple results](#multiple-results)
	- [Named return values](#named-return-values)
	- [Short variable declarations](#short-variable-declarations)
	- [Zero values](#zero-values)
	- [Constants](#constants)
- [Flow control statements: for, if, else, switch and defer](#flow-control-statements-for-if-else-switch-and-defer)
	- [For](#for)
	- [If](#if)
	- [Switch](#switch)
	- [Defer](#defer)
	- [Stacking defers](#stacking-defers)
- [More types: structs, slices, and maps.](#more-types-structs-slices-and-maps)
	- [Pointers](#pointers)
	- [Structs](#structs)
	- [Pointers to structs](#pointers-to-structs)
	- [Arrays](#arrays)
	- [Slices](#slices)
	- [Slices are like references to arrays](#slices-are-like-references-to-arrays)
	- [Slice length and capacity](#slice-length-and-capacity)
	- [Nil slices](#nil-slices)
	- [Creating a slice with make](#creating-a-slice-with-make)
	- [Appending to a slice](#appending-to-a-slice)
	- [Range](#range)
	- [Range continued](#range-continued)
	- [Maps](#maps)
	- [Mutating Maps](#mutating-maps)
	- [Function values](#function-values)
	- [Function closures](#function-closures)
- [Methods and interfaces](#methods-and-interfaces)
	- [Methods](#methods)
	- [Methods continued](#methods-continued)
	- [Pointer receivers](#pointer-receivers)
	- [Interfaces](#interfaces)
	- [Interface values](#interface-values)
	- [Interface values with nil underlying values](#interface-values-with-nil-underlying-values)
	- [Stringers](#stringers)
	- [Errors](#errors)
- [Generics](#generics)
	- [Type parameters](#type-parameters)
	- [Generic types](#generic-types)
- [Concurrency](#concurrency)
	- [Goroutines](#goroutines)
	- [Channels](#channels)
	- [Buffered Channels](#buffered-channels)
	- [Range and Close](#range-and-close)
	- [Select](#select)
	- [Default Selection](#default-selection)
	- [sync.Mutex](#syncmutex)

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

ポイント：
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

ポイント：
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

ポイント：
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

ポイント：
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

ポイント：
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

実行結果：
```
55
```

ポイント：
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

ポイント：
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

ポイント：
- Goでは戻り値となる変数に名前をつけることができる
- 名前付きの戻り値の変数を使うと、returnで何も書かずに戻せる
- 短い関数以外では使わない方が良い


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

ポイント：
- Goでは`var i int = 10`という書き方で変数宣言もできる
- `i := 10`という代入文を使って暗黙的な型宣言もできる
- 関数内では`:=`を積極的に使うべき
- 関数外と、明示的に型指定したい場合に`var`を使う(int64、float64、構造体ゼロ値など)


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

ポイント：
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

### Constants

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

実行結果：
```
Hello 世界
Happy 3.14 Day
Go rules? true
```

ポイント：
- const は 変わらない値 を表す
- := では宣言できない（const x = 1）
- 使える型は 数値・文字列・bool のみ
- Goの const は「値」そのものを表し、型に縛られない
  - 使われる場所（文脈）で、最適な型に“はめ込まれる”
  - 型なし定数（untyped constant）と呼ばれる
    ```go
    const n = 10
    var a int     = n  // OK
    var b int64   = n  // OK
    var c float64 = n  // OK
    ```
    - n は代入先の型として解釈される（文脈で型が決まる）
    - math / time / DB で効果的
      - (例)math（r が変数だと、型合わせが必要。）
      ```
      const r = 3
      area := math.Pi * r * r  // OK
      ```
- マジックナンバーを消すためにも使う
- 設計上の 「不変」という意図を表現できる

---

## Flow control statements: for, if, else, switch and defer

### For

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

実行結果：
```
45
```

ポイント：
- Go のループ構文は `for` だけ（`while` / `do-while` はない）
- 基本形は `for 初期化; 条件; 後処理 { }`
- `()` は不要、`{}` は必須
- `for i := 0; ...` の変数は for のスコープ内のみ有効
- 条件が false になるとループ終了
- 無限ループは `for {}` で書く

---

### If

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

実行結果：
```
1.4142135623730951 2i

```

ポイント：
- IFは `if 条件 { }` が基本形。`()` は不要、`{}` は必須
- 条件式は bool 型のみ（数値や nil は不可）
- if の前で 短縮変数宣言ができる
    ```go
    if v := f(); v > 0 { }
    ```
- その変数のスコープは if / else 内のみ
- else if / else が使える
- else は 必ず直前の if の } と同じ行

---

### Switch

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

実行結果：
```
Go runs on Linux.
```

ポイント：
- `switch` は `if-else` の簡潔な書き方
- 他の言語と異なり、自動で `break` する（次のcaseに行かない）
  - `fallthrough`構文はあるが、次の case の条件を評価せずに本体を実行するため注意が必要
  - 2つの条件を両方評価して実行したいならswitchは使わずifを使う
- case は 定数でなくてよい
- 式の型は自由（int 以外もOK）
- switch で 短縮変数宣言ができる（if同様）
- `switch {}` で if 代わりに使える（if-else より 読みやすく条件分岐の「列挙」に向いている）
- default は省略可

---

### Defer

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

実行結果：
```
hello
world
```

ポイント：
- defer は、指定した関数の実行を呼び出し元の関数終了直前まで遅延させる
- defer に渡した 関数の引数は、その場ですぐ評価される
  - 以下例がわかりやすい
    ```go
    func main() {
      x := "world"
      defer fmt.Println(x) // world
      x = "golang"
    }
    ```
  - `defer fmt.Println(x)` の時点で x の値は `"world"`
  - その値が `defer` 用に保存される
  - 後で `x` が変わっても関係ない
- 実際の関数呼び出しは return 直前（panic 時も含む） に行われる
- defer された処理は、後に書いたものから先に実行される（LIFO 後入れ先出し）
- 関数内で `return` がどこにあっても、`defer` は必ず実行される
- 主な用途は リソースの後片付け（ファイルClose, mutexのUnlock など）
- `panic` が起きても実行される（`recover` があれば復旧できる）
- Go らしい書き方
  - 成功直後に defer を書く 
  - 「後始末がある処理の近くに書く」ことで可読性が上がる
    ```go
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    // 以降は安心して処理を書ける
    ```
- Go では エラー処理や安全な終了を簡潔に書くための重要な構文

---

### Stacking defers

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

実行結果：
```
counting
done
9
8
7
6
5
4
3
2
1
```

ポイント：
- このコードの実行順序
  - 1.`fmt.Println("counting")` → `counting`
  - 2.for ループで `i = 0..9` の defer が積まれる
    - i=0 → defer fmt.Println(0)
    - i=1 → defer fmt.Println(1)
    - …
    - i=9 → defer fmt.Println(9)
  - 3.`fmt.Println("done")` → `done`
  - 4.関数終了 → defer が LIFO(後入れ先出し) で実行
    - fmt.Println(9)
    - fmt.Println(8)
    - …
    - fmt.Println(0)
- 複数の defer は スタックに積まれる（LIFO(last-in-first-out) ：後入れ先出し）
- 関数終了時に、最後に書いた defer から順に実行される
- 引数は defer 行で即評価される
- for ループで defer すると 逆順で実行される
- 複数リソースや mutex の解放順を自然に制御できる
- panic が起きても defer は LIFO 順で実行される
- 高頻度ループでの defer は メモリコストに注意
- クロージャは外側変数を参照する（**わかりづらいので後でまとめ直したい**）
  - そのため、for ループ変数を無名関数で defer すると、最後の値しか参照されない
  - 意図した値を出すには 変数をコピーする か クロージャに引数として渡す

---

## More types: structs, slices, and maps.

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

ポイント：
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

ポイント：
- struct (構造体)は、フィールドの集まり
- `type 名前 =? 型定義` の構文で型に名前をつける。構造体以外でも使う。
  - 構造体：`type A struct {}`
  - インターフェース：`type B interface {}`
  - 独自型：`type UserID int`
  - 関数：`type Fn func()`
  - スライス/マップ：`type IDs []int`


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

ポイント：
- structのフィールドは、structのポインタを通してアクセスすることもできる
- `p.X`でアクセスできる
   - ちなみに`(*p).X`と書くことでアクセスもできるが長いのであまり使われない

---

### Arrays

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

実行結果：
```
Hello World
[Hello World]
[2 3 5 7 11 13]
```

ポイント：
- Goの配列は固定長
- [n]T は型 T の長さ n の配列を表す
- 配列の長さは型の一部で、後から変更できない
- インデックスで要素にアクセス・代入可能 (a[0] = "Hello")
- 配列リテラルで初期化可能 (primes := [6]int{2,3,5,7,11,13})
- 配列をまとめて出力できる (fmt.Println(a))

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

ポイント：
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

ポイント：
- スライスは配列そのものを持たず、配列の一部を“参照”している
- スライス同士が同じ配列を共有することがある
- どれか1つのスライスを変更すると、配列・他のスライスにも影響する
- コピーしたい場合は`copy()`か`append()`を使う


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

ポイント：
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


ポイント：
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

ポイント：
- make([]T, len, cap) は 配列を確保し、その配列を参照するスライスを作る
- len は「今使っている長さ」、cap は「将来使える余白」
- cap を指定することで append 時の再確保を防げる
- 高頻度処理・件数が見積もれる場合は make を使う意味がある
- make([]T, 0, n) は「これから詰める」意図を明示できる
- `[:n]` →「先頭から n 個見る」
- `[i:j]` →「i から j-1 まで見る」
- パッと見て分かりづらいのでまとめ直したい


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

ポイント：
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

ポイント：
- Goの`for + range`は、PHP等の`foreach`とほぼ同じ
- `for + range`は値コピーなので元の値(pow)が書き換えられない
  - PHPもデフォルトは値コピー。オプションで参照にもできるが基本使わないのでほぼ同じと考えてOK


---

### Range continued

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

実行結果：
```
1
2
4
8
16
32
64
128
256
512
```

ポイント：
- `for i, _ := range slice` → 値を捨てる
- `for _, value := range slice` → インデックスを捨てる
- `for i := range slice` → 値を使わずインデックスだけ使う
- `_` は「使わない変数」を表す特殊識別子



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

ポイント：
- Goの連想配列は`map`
- mapのゼロ値は `nil`
- PHPと異なり、`make`しないと要素追加できない(宣言だけでは使えない)
- Goのmapは型安全なのでキーごとに異なる型の値は入れられない(PHPは入れられる)


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

ポイント：
- mapは 存在しないキーでもゼロ値が返る
- 存在確認が必要な場合は必ず ok を使う
  - valueはゼロ値が入るため存在確認に使えない
- 追加・更新・削除は安全で、存在しないキーに対する操作でもpanicしない


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

ポイント：
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

ポイント：
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

ポイント：
- Goにクラスはない
- 型に対してメソッドを定義できる（構造体以外にも独自型にも定義可能）
- このコードでは `(v Vertex)` の部分がレシーバ
- メソッドは `v.Abs()` のように呼べる
- 「データ + 振る舞い」を表現するための仕組み

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

実行結果：
```
1.4142135623730951
```


ポイント：
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

ポイント：
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

ポイント：
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

ポイント：
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

ポイント：
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

### Stringers

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

実行結果：
```
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

ポイント：
- Stringer は fmt が表示用に使う超重要インターフェース
- String() string を実装するだけで自動的に使われる
- fmt.Println / fmt.Printf("%v") は Stringer を優先して表示する
- 表示ロジックを型に閉じ込められる
- ログ・デバッグ・ドメイン表現がきれいになる
- Stringerで加工された表示ではなくプレーンな内部構造を見たいときは fmt.Printf の指定子（`%+v` `%#v`等）で切り替える

---

### Errors

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

実行結果：
```
at 2026-01-09 11:54:35.282602 +0900 JST m=+0.000180501, it didn't work
```

ポイント：
- error は Stringerと同様にただのインタフェース
- 独自エラー型は `struct + Error()` で作る
- Goでのエラーハンドリングは、例外（try/catch）を使わず、「値としてエラーを返す」 という設計
  - `err != nil` が失敗 / `err == nil` が成功
- fmt は Error() を自動で呼ぶ（error → Stringer → デフォルト表示 の順）

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

ポイント：
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

### Generic types

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

ポイント：
- ジェネリクスは 関数だけでなく構造体（struct）にも使える
- T は struct 全体で共有される型パラメータ
- フィールドにも T をそのまま使える
  - 値保持用：val T
  - 再帰参照：*List[T]
- 上のコードでは型安全な再帰構造を作っている
- 実務では自作することは少ない

---

## Concurrency

### Goroutines

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

実行結果：（出力順は毎回変わる）
```
hello
world
hello
world
world
hello
hello
world
hello
```

ポイント：
- goroutine（ゴルーチン）は、Goのランタイムに管理される軽量な実行単位
  - OSスレッドそのものではない
  - Goランタイムが必要に応じてOSスレッドに割り当てる（goroutine は OSスレッドの上に乗っている）
  - OSスレッド より 圧倒的に軽い
  - 数万〜数十万起動しても現実的
  - `Goが用意しためちゃくちゃ軽いスレッドのようなもの` くらいの理解でOK
- `main()` もgoroutineの一つ。一番最初に実行されるgoroutine
    ```
    func main() {
    	go say("world") // 新しいgoroutine で実行される
    	say("hello")    // main goroutine で実行される
    }
    ```
  - すべてのコードは必ず何らかの goroutine 上で動く
  - main goroutine が終了すると、プログラム全体が終了する
- `go say("world")` の意味は「関数呼び出し」ではなく「並行実行の開始」
  - ① `main goroutine`で引数を準備する
  - ② `新しいgoroutine`を起動
  - ③ sayの処理は`新しいgoroutine`に任せる
  - ④ 自分(`main goroutine`)は止まらず次へ進む
- 実行結果の出力順は保証されない
  - スケジューリングは Go ランタイム任せで、実行タイミングは毎回微妙に異なる
- goroutine は同じアドレス空間で実行される
  - メモリは共有されるため、データ競合の危険がある
  - Goでは次章以降で解説する channel を使ったデータのやり取りが推奨される

---

### Channels

```go
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
```

実行結果：（順不同）
```
-5 17 12
```

ポイント：
- チャネルは goroutine 間で値を受け渡す通り道
- `ch <- v` // v をチャネル ch へ送信する
- `v := <-ch` // ch から受信した変数を v へ割り当てる
- 非バッファチャネルでは
  - `ch <- v` は 受信側が来るまでブロック
  - `<-ch` は 送信側が来るまでブロック
  - 送信と受信は 同時に揃った瞬間だけ値が渡る
  - この仕組みは `握手モデル` と呼ばれる
- ブロックにより：
  - goroutine の 同期(完了待ち) が自然に書ける
  - ロックや条件変数が不要
  - 値が中途半端な状態で観測されない
  - ブロックは 無駄な待ちではない
    - OSスレッドを止めない
    - Goランタイムが安全にスケジューリング
    - <-ch は Go における await 相当
- goroutine で動かす前提の関数は、戻り値の代わりにチャネルへ送信するのが基本
- チャネルは データ転送＋同期 を同時に実現する

---

### Buffered Channels

```go
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
```

実行結果：
```
1
2
3
```

ポイント：
- バッファ付きチャネルは 値を一時的に溜めるキュー
- `ch := make(chan int, 2)` でmake の２つ目の引数にバッファの長さを与える
- 送信（ch <- v）のルール
  - バッファに 空きがあれば → すぐ送信できる（ブロックしない）
  - バッファが 満杯なら → 受信されるまでブロック
- 受信（<-ch）のルール
  - バッファに 値があれば → すぐ受信できる
  - バッファが 空なら → 送信されるまでブロック
- デッドロックは、送信や受信で処理が止まったとき(ブロックした時)、その待ち状態を解消できる goroutine が他に存在しない場合にデッドロックが起きる。

---

### Range and Close

```go
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

実行結果：
```

```

ポイント：
- close(ch) は「これ以上値を送らない」という終了通知
- チャネルを close しても受信は可能（送信は不可・panicになる）
- close すべきなのは送信側。受信側が close すると設計破綻しやすい
- 受信時は `v, ok := <-ch` で close を検知できる
  - `ok == false` → チャネルが close され、値がない
- `for v := range ch` を使えば自分でokの判定を書かなくて良い。
  - 内部的に以下と同じことをしている
    ```
    for {
      v, ok := <-ch
      if !ok {
          break
      }
      ...
    }
    ```
  - `range ch` はチャネルが close されるまで受信し続ける
  - チャネルが close され、かつ空になった瞬間に自動で終了
  - close がないと `range ch` は終了せず、デッドロックする
- チャネルは **基本 close しなくてよい**
  - 単発の送受信やリクエスト/レスポンス型では不要
  - close が必要なのは「もう値が来ないことを受け手に伝えたいとき」だけ
    - 例えば、for range、ストリーム処理、ワーカー終了通知


---

### Select

```go
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
```

実行結果：
```
0
1
1
2
3
5
8
13
21
34
quit
```

ポイント：
- 各チャネルの役割（ややこしいので先に整理）
  - `c`チャネル は値（データ）を流すためのチャネル
    - 送信側：fibonacci
    - 受信側：無名関数（go func）
  - `quit`チャネル は制御（終了通知）を流すためのチャネル
    - 送信側：無名関数（go func）
    - 受信側：fibonacci
- このコードの流れ
  - fibonacci が c <- x を試みる（送信する）
  - 無名関数 が <-c する（受信する）
  - フィボナッチ数が表示される
  - これを 10 回繰り返す
  - 無名関数のforループが終了して quit <- 0（終了通知を送信）
  - select の case <-quit が成立（終了通知を受信）
  - "quit" を表示して return（fibonacci関数終了）
  - プログラム終了
- `select`とは
  - select は 複数のチャネル操作を同時に待つための構文
  - select は、準備できた case を1つ実行する。複数準備できていれば ランダムに選択
- `quit`とは
  - quit チャネルは「処理をやめろ」という 制御シグナル
  - select と組み合わせて使う
  - goroutine が 自分で return して終了
- `close`と`quit`の役割の違い
  - close → データの終端を表す
  - quit → 処理の中断・キャンセルを表す
- 設計指針
  - 値の列が自然に終わる → close
  - 途中で止めたい／外部から制御したい → quit / context
- 実務では `quit` はほぼ `ctx.Done()` に置き換わる
- select + quit を理解できるとcontext.Context の ctx.Done() が理解できる
- `fibonacci()`の引数について
  - 実務ではチャネルの方向性を明示した書き方がベター
    ```go
    func fibonacci(c chan<- int, quit <-chan int) {
    ```
  - chan<- int：送信専用（c）
  - <-chan int：受信専用（quit）
  - 「役割が固定されている」ことを型で保証できる
- まとめると
  - close は「データ終了」
  - quit は「処理中断」

---

### Default Selection

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // 指定間隔ごとに通知するチャネル
	boom := time.After(500 * time.Millisecond) // 指定時間後に1回だけ通知するチャネル
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

実行結果：
```
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
BOOM!
```

ポイント：
- select は 複数のチャネル操作を同時に待つ制御構文
- 実行可能な case が1つでもあれば、それが選ばれる
- default を書くと
  - どの case も準備できていない場合に即実行
  - select が ブロックしなくなる
- default 付き select は
  - ポーリング / ゲームループ / 監視処理 に使われる
  - sleep なしだと CPU を浪費する
- time.Tick / time.After は「時間イベントを流すチャネル」

---

### sync.Mutex

```go
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
```

実行結果：
```
1000
```

ポイント：
- goroutine は軽量で簡単に並行実行できるが、共有データへの同時アクセスは自動では安全（ゴルーチンセーフ）にならない
- goroutine 間で 共有変数を安全に扱うために排他制御が必要
- 排他制御（mutual exclusion）を実現する仕組みが Mutex
- Goでは `sync.Mutex` を使い、`Lock()` / `Unlock()` で制御する
- 同時に1つの goroutine だけが Lock 区間を実行できる
- map など ゴルーチンセーフ でないデータは Mutex で保護する
- Lock と Unlock は必ずペアで使う
- `defer mu.Unlock()` を使うと Unlock 漏れを防げる
- 読み取り（参照）だけでも Lock が必要
- goroutine 間で「通知・受け渡し・停止」を表現したい場合はチャネルが向くが、単に共有データを安全に守りたい場合は Mutex が最適

---

以上
