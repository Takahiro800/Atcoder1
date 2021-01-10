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
	A [][]int
)

func main() {
	N = ReadInt()
	n := pow(2, N)
	A = make([][]int, N)
	A[0] = ReadIntSlice(n)

	for i := 0; i < N-1; i++ {
		A[i+1] = buttle(A[i], n)
		n /= 2
	}

	x := A[N-1][0]
	y := A[N-1][1]

	sec := smallInt(x, y)

	for i := 0; i < pow(2, N); i++ {
		if sec == A[0][i] {
			fmt.Println(i + 1)
			return
		}
	}
}

func smallInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func pow(a, b int) int {
	ans := 1
	for i := 0; i < b; i++ {
		ans *= a
	}
	return ans
}

func buttle(A []int, s int) []int {
	s = s / 2

	B := make([]int, s)
	for i := 0; i < s; i++ {
		B[i] = largeInt(A[2*i], A[2*i+1])
	}
	A = B
	return A
}

func largeInt(a, b int) int {
	if a > b {
		return a
	}
	return b
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
