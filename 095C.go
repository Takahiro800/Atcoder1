package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INF_BIT60 = 1 << 60
)

func main() {
	b := ReadIntSlice(5)
	A := b[0]
	B := b[1]
	C := b[2]

	X := b[3]
	Y := b[4]
	ans := INF_BIT60

	for k := 0; k <= max(X, Y); k++ {
		pay := A*max(0, X-k) + B*max(0, Y-k) + 2*C*k
		// pay := A*i + B*j + C*k
		if ans > pay {
			ans = pay
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
