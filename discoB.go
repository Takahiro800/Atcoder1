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
	A []int
)

func main() {
	N = ReadInt()
	A = ReadIntSlice(N)
	sum := 0

	for _, v := range A {
		sum += v
	}

	left := 0
	right := sum
	min_diff := sum

	for _, v := range A {
		left += v
		right -= v
		diff := abs(left - right)
		if min_diff > diff {
			min_diff = diff
		}
	}

	fmt.Println(min_diff)

}

// snipet

// INF_BIT60 = 1 << 60
var sc = bufio.NewScanner(os.Stdin)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

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