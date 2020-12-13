package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N, D int
	X    [][]float64
)

func main() {
	N, D = ReadInt(), ReadInt()
	X = make([][]float64, N)
	ans := 0

	for i := 0; i < N; i++ {
		X[i] = ReadFloatSlice(D)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if judgeInt(distance(X[i], X[j])) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func distance(a, b []float64) float64 {
	x := 0.0
	for i := 0; i < len(a); i++ {
		x += (a[i] - b[i]) * (a[i] - b[i])
	}
	return math.Sqrt(x)
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func ReadFloatSlice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat()
	}
	return b
}

func ReadFloat() float64 {
	v, e := strconv.ParseFloat(Read(), 64)
	if e != nil {
		panic(e.Error())
	}
	return v
}

func judgeInt(x float64) bool {
	if math.Floor(x) == x {
		return true
	}
	return false
}

// MergeSortについて
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(left, right []int) (ret []int) {
	ret = []int{}
	for len(left) > 0 && len(right) > 0 {
		var x int
		// ソート済みのふたつのスライスからより小さいものを選んで追加していく (これがソート処理)
		if right[0] > left[0] {
			x, left = left[0], left[1:]
		} else {
			x, right = right[0], right[1:]
		}
		ret = append(ret, x)
	}
	// 片方のスライスから追加する要素がなくなったら残りは単純に連結できる (各スライスは既にソートされているため)
	ret = append(ret, left...)
	ret = append(ret, right...)
	return
}

func sort(left, right []int) (ret []int) {
	// ふたつのスライスをそれぞれ再帰的にソートする
	if len(left) > 1 {
		l, r := split(left)
		left = sort(l, r)
	}
	if len(right) > 1 {
		l, r := split(right)
		right = sort(l, r)
	}

	// ソート済みのふたつのスライスをひとつにマージする
	ret = merge(left, right)
	return
}

func split(values []int) (left, right []int) {
	// スライスを真ん中でふたつに分割する
	left = values[:len(values)/2]
	right = values[len(values)/2:]
	return
}

func Sort(values []int) (ret []int) {
	left, right := split(values)
	ret = sort(left, right)
	return
}


