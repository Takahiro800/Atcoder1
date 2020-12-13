package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	N int
	W []string
)

func main() {
	N = ReadInt()
	W = make([]string, N)

	W[0] = Read()
	for i := 1; i < N; i++ {
		w := Read()
		if checkUniqueness(w, W) {
			W[i] = w
		}
	}

	for i := 1; i < N; i++ {
		w := W[i]
		if checkContinuity(w, W[i-1]) {
			W[i] = w
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}

func checkUniqueness(target string, list []string) bool {
	for i, _ := range list {
		if list[i] == target {
			return false
		}
	}
	return true
}

func checkContinuity(target, prev string) bool {
	return strings.HasSuffix(prev, returnString(target, 0))
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

func returnString(n string, i int) string {
	return string([]rune(n)[i])
}
