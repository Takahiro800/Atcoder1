package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	S string
	T string
)

func main() {
	S = Read()
	T = strings.Repeat("01", 5*10*10*10*10*10*10*10)
	l := len(S)
	a := T[0:l]
	b := T[1 : l+1]
	k := countDif(a, S)
	t := countDif(b, S)

	ans := min(k, t)
	fmt.Println(ans)
}

func countDif(x, y string) int {
	ans := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			ans++
		}
	}
	return ans
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
