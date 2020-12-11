package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N int
	H []int
)

func main() {
	N = ReadInt()
	H = ReadIntSlice(N)
	ans := 0
	a := 0
	for i := 0; i < N-1; i++ {
		if H[i] < H[i+1] {
			if ans < a {
				ans = a
			}
			a = 0
			continue
		}
		a++
	}
	if a > ans {
		ans = a
	}
	fmt.Println(ans)
}

func checkHeight(c []int, n int) int {
	a := 0
	for i := n; i < len(c)-1; i++ {
		if c[i] < c[i+1] {
			return a
		}
		a++
	}
	return a
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
