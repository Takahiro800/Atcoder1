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
	S, T []string
)

func main() {
	N = ReadInt()
	Q := make([]string, N)
	for i := 0; i < N; i++ {
		Q[i] = Read()
	}
	S = make([]string, N)
	T = make([]string, N)

	for i := 0; i < N; i++ {
		j := 0
		k := 0
		x := Q[i]
		if x[0] == 33 {
			for m := 0; m < N; m++ {
				if "!"+S[m] == x {
					fmt.Println(S[m])
					return
				}
			}
			T[k] = x
			k++
		} else {
			if contains(T, "!"+x) {
				fmt.Println(x)
				return
			} else {
				S[j] = x
				j++
			}
		}
	}
	fmt.Println("satisfiable")
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
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
