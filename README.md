
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
### [未]Short variable declarations
### [未]Basic types
### [未]Zero values
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

### [未]Pointers
### [未]Structs
### [未]Struct Fields
### [未]Pointers to structs
### [未]Struct Literals
### [未]Arrays
### [未]Slices
### [未]Slices are like references to arrays
### [未]Slice literals
### [未]Slice defaults
### [未]Slice length and capacity
### [未]Nil slices
### [未]Creating a slice with make
### [未]Slices of slices
### [未]Appending to a slice
### [未]Range
### [未]Range continued
### [未]Exercise: Slices
### [未]Maps
### [未]Map literals
### [未]Map literals continued
### [未]Mutating Maps
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
