# Atcoder1
- `taekari`のAtcoderのコードと途中から自分の所感、勉強したことをコメントとして残しています。
- 2020/3/24　現在レート106のよわよわエンジニアです。
- 2020/12/02 Atcoder Problems再開
- 当分はGoで取り組みます
- 年度内に着色します

# Goの文法メモ
## 同じ文字列を繰り返して連結する
```go
str := "abe"
strings.Repeat(str, 3)
```
## 指定した桁の数字を取得したい
- byte型になる
- 1の値が49なので、48を引いてあげれば良い
- int型に変換する場合は `int()`
```go
s string
a byte

num = a - 48
numInt := int(num)
```

## 文字列中の特定の文字をカウントしたい
```go
strings.Count(文字列, 特定の文字 string ) int
```