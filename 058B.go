package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	O, E, ans string
)

func main() {
	O = Read()
	E = Read()
	if len(O) == len(E) {
		for i := 0; i < len(O); i++ {
			ans += string([]rune(O)[i]) + string([]rune(E)[i])
		}
	} else {
		ans += string([]rune(O)[0])
		for i := 0; i < len(E); i++ {
			ans += string([]rune(E)[i]) + string([]rune(O)[i+1])
		}
	}
	fmt.Println(ans)

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
