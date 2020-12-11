package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, M int
	S    [][]int
	P    []int
)

func main() {
	b := ReadIntSlice(2)
	N = b[0]
	M = b[1]

	for i := 0; i < M; i++ {
		k := ReadInt()
		row := ReadIntSlice(k)

		for j := 0; j < len(row); j++ {
			row[j]--
		}
		S = append(S, row)
	}
	P = ReadIntSlice(M)

	ans := 0
	for bits := 0; bits < (1 << uint64(N)); bits++ {
		flg := true
		// bitsの各要素についてのループ
		for i := 0; i < M; i++ {
			c := 0
			for _, j := range S[i] {
				// bitsのi個目の要素の状態がonかどうかチェック
				if (bits>>uint64(j))&1 == 1 {
					// 何かしらの処理
					c++
				}
			}
			if c%2 != P[i] {
				flg = false
				break
			}
		}
		if flg {
			ans++
		}
	}
	fmt.Println(ans)

}

var sc = bufio.NewScanner(os.Stdin)

func input() string {
	sc.Scan()
	return sc.Text()
}

func init()        { sc.Buffer(make([]byte, 256), 1e9); sc.Split(bufio.ScanWords) }
func Read() string { sc.Scan(); return sc.Text() }

func ReadInt() int {
	v, e := strconv.Atoi(Read())
	if e != nil {
		panic(e.Error())
	}
	return v
}

func ReadString() string {
	if !sc.Scan() {
		panic(sc.Err())
	}
	return sc.Text()
}

// ReadInstSlice returns an integer slice that han n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}
