# Go学習メモ
## 標準入力

```go
var x, n int
fmt.Scan(&x, &n)
```

## 文字列を整数値に変換
[`strconv.Atoi`](https://golang.org/pkg/strconv/#Atoi)を使いましょう

```go
func main() {
  v := "10"
  if s, err := strconv.Atoi(v); err == nil {
    fmt.Printf("%T, %v", s, s)
  }
}
```
## 書式を指定して標準出力に書き込みを行う Printf

- `%T`
  - 対象のデータの情報を埋め込む
- `%v`
  - デフォルトフォーマットで対象のデータの情報を埋め込む
- `%t`
  - 論理値
- `%d`
  - 符号付き整数
- `%d`
  - 符号なし整数
