package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ans := 0
	count := 0
	b := ReadIntSlice(3)
	A := b[0]
	B := b[1]
	K := b[2]
	m := min(A, B)
	for i := m; i >= 1; i-- {
		if A%i == 0 && B%i == 0 {
			count++
			if count == K {
				ans = i
				break
			}
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
