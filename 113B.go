package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N    int
	T, A float64
	H    []float64
)

func main() {
	N = ReadInt()
	T = ReadFloat()
	A = ReadFloat()
	H = ReadFloatSlice(N)

	dif := 10000.0
	ans := 0
	for i := 0; i < N; i++ {
		t := calcTemperature(T, H[i])
		d := absFloat(A - t)
		if d < dif {
			dif = d
			ans = i + 1
		}
	}
	fmt.Println(ans)
}

func calcTemperature(T, x float64) float64 {
	t := T - x*0.006
	return t
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

func absFloat(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -1 * a
}
