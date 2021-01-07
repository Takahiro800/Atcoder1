package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	x, y int
)

func main() {
	x = ReadInt()
	y = ReadInt()

	if group(x) == group(y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func group(i int) int {
	for _, a := range []int{1, 3, 5, 7, 8, 10, 12} {
		if a == i {
			return 1
		}
	}
	for _, a := range []int{4, 6, 9, 1} {
		if a == i {
			return 2
		}
	}
	return 3
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
