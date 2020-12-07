package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ReadInt()
	ans := 12
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			a := N / i
			st := strconv.Itoa(a)
			if ans > len(st) {
				ans = len(st)
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
