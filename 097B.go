package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	x int
)

func main() {
	x = ReadInt()
	A := 1
	for n := 1; n < 10; n++ {
		if x < double(n) {
			break
		} else {
			A = double(n)
		}
	}
	B := 1
	for b := 3; b < 32; b++ {
		if x < mult(b, 2) {
			break
		} else {
			B = mult(b, 2)
		}
	}
	C := 1
	for b := 3; b < 11; b++ {
		if x < mult(b, 3) {
			break
		} else {
			C = mult(b, 3)
		}
	}
	D := 1
	for b := 3; b < 6; b++ {
		if x < mult(b, 4) {
			break
		} else {
			D = mult(b, 4)
		}
	}
	E := 1
	for i := 5; i < 7; i++ {
		if x < mult(3, i) {
			break
		} else {
			E = mult(3, i)
		}
	}
	fmt.Println(Max(A, B, C, D, E))
}

func double(n int) int {
	a := 1
	for i := 0; i < n; i++ {
		a *= 2
	}
	return a
}

func mult(b, p int) int {
	a := 1
	for i := 0; i < p; i++ {
		a *= b
	}
	return a
}

func Max(s ...int) int {
	a := 1
	for _, v := range s {
		if a < v {
			a = v
		}
	}
	return a
}

// snipet
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
