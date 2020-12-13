package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := ReadIntSlice(2)
	mergeSort(S)
	fmt.Println(S)
}

func merge(a, b []int) (S []int) {
	S = []int{}

	for len(a) > 0 && len(b) > 0 {
		var s int
		if a[0] >= b[0] {
			s, a = a[0], a[1:]
		} else {
			s, b = b[0], b[1:]
		}
		S = append(S, s)
	}
	S = append(S, a...)
	S = append(S, b...)
	return
}

func mergeSort(S []int) []int {
	left := S[:len(S)/2]
	right := S[len(S)/2:]
	mergeSort(left)
	mergeSort(right)
	S = merge(left, right)
	return S
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
