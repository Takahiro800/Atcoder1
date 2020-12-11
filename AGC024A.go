package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	A, B, C, K int
)

func main() {
	in := ReadIntSlice(4)
	A = in[0]
	B = in[1]
	C = in[2]
	K = in[3]
	var code int

	if K%2 == 1 {
		code = -1
	} else {
		code = 1
	}
	ans := code * (A - B)

	fmt.Println(ans)
}

func otherSum(b, c int) int {
	return b + c
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

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
