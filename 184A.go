package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	a, b, c, d int
)

func main() {
	A := ReadIntSlice(2)
	a = A[0]
	b = A[1]

	B := ReadIntSlice(2)
	c = B[0]
	d = B[1]

	ans := a*d - b*c
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
	z := make([]int, n)
	for i := 0; i < n; i++ {
		z[i] = ReadInt()
	}
	return z
}
