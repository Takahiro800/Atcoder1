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

## スライスをsortする
`sort.Sort(sort.IntSlice(v))`でint型のスライスをソートする



## 構造体をsortする
`sort.Slice("構造体のスライス", func(i,j int) bool { 比較条件 })`　でソートできる

```go
type shop struct {
  name string
  point int
}

func sort_shop {
  sort.Slice(shop_list, func(i,j int) bool {
    if shop_list[i].name == shop_list[j].name {
          shop_list[i].point > shop_list[j].point
    } else {
          shop_list[i].name <shop_list[j].name
    }
  })
}
```

## とりあえずスライスを作成して、追加していく
```go
//長さ０のスライスを用意
v := make([]int, 0)

//末尾に追加する
v = append(v, element)
```