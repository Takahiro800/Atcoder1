package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	S   string
	K   int
	a   byte
	ans byte
)

func main() {
	S = Read()
	n := len(S)
	K = ReadInt()
	ans = 1
	num := 1

	for i := 0; i < n; i++ {
		a = S[i] - 48

		if a != 1 {
			ans = a
			num = i + 1
			break
		}
	}
	// fmt.Println(K, num, ans)
	if K < num {
		fmt.Println(1)
	} else {
		fmt.Println(ans)
	}

	fmt.Printf("%v %T\n", int(a), int(a))
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
