package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	n, m int
	s, t map[string]int
)

func main() {
	n = ReadInt()
	s = make(map[string]int)
	for i := 0; i < n; i++ {
		key := Read()
		if _, ok := s[key]; ok {
			s[key]++
		} else {
			s[key] = 1
		}
	}
	m = ReadInt()
	for i := 0; i < m; i++ {
		key := Read()
		if _, ok := s[key]; ok {
			s[key]--
		} else {
			s[key] = -1
		}
	}
	sum := 0
	for _, v := range s {
		if sum < v {
			sum = v
		}
	}
	fmt.Println(sum)
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
