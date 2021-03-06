package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N int
)

type button struct {
	next int
	flag bool
}

func main() {
	N = ReadInt()
	button_list := make([]button, N)
	for i := 0; i < N; i++ {
		button_list[i].next = ReadInt()
		button_list[i].flag = false
	}

	solve(button_list, 1, 1)
}

func solve(A []button, hoge, ans int) {
	a := A[hoge-1].next
	if A[a-1].flag == true {
		fmt.Println(-1)
		return
	} else if a == 2 {
		fmt.Println(ans)
		return
	} else {
		ans += 1
		A[hoge-1].flag = true
		solve(A, a, ans)
	}
}

// snipet

// INF_BIT60 = 1 << 60
var sc = bufio.NewScanner(os.Stdin)

func input() string {
	sc.Scan()
	return sc.Text()
}

func init() { sc.Buffer(make([]byte, 256), 1e9); sc.Split(bufio.ScanWords) }

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

func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
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
