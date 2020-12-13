package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	A, B, K int
)

func main() {
	A, B, K = ReadInt(), ReadInt(), ReadInt()

	if 2*K > B-A+1 {
		for i := A; i <= B; i++ {
			fmt.Println(i)
		}
	} else {
		for i := 0; i < K; i++ {
			fmt.Println(A)
			A++
		}
		B = B - K + 1
		for i := 0; i < K; i++ {
			fmt.Println(B)
			B++
		}
	}
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
