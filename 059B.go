package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	A, B string
)

func main() {
	A = Read()
	B = Read()
	a, _ := strconv.Atoi(A)
	b, _ := strconv.Atoi(B)
	if a > b {
		fmt.Println("GREATER")
	} else if a < b {
		fmt.Println("LESS")
	} else {
		fmt.Println("EQUAL")
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
