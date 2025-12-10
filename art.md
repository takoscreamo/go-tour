
実務でGoを扱うので、改めて [A Tour of Go](https://go-tour-jp.appspot.com/) を写経しながら、自分用メモとして更新していきます。

## Hello, World

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


## Packages

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

## Imports

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

## Exported names

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

## Functions

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


