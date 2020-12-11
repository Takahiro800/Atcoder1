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
	S string
)

func main() {
	N = ReadInt()
	S = Read()
	ans := 0
	for i := 1; i < 10; i++ {
		s := strconv.Itoa(i)
		a := strings.Index(S, "0")
		if a >= 0 {
			b := strings.Index(S[a+1:], "0")
			if b > 0 {
				c := strings.Index(S[b+1:], string(s))
				if c > 0 {
					ans++
					fmt.Println(i)
				}
			}
		}
	}
	for i := 10; i < 100; i++ {
		s := strconv.Itoa(i)
		a := strings.Index(S, "0")
		if a >= 0 {
			b := strings.Index(S[a+1:], string(s[0]))
			if b > 0 {
				c := strings.Index(S[b+1:], string(s[1]))
				if c > 0 {
					ans++
					fmt.Println(i)
				}
			}
		}
	}
	for i := 100; i < 1000; i++ {
		s := strconv.Itoa(i)
		a := strings.Index(S, string(s[0]))
		if a >= 0 {
			b := strings.Index(S[a+1:], string(s[1]))
			if b > 0 {
				c := strings.Index(S[b+1:], string(s[2]))
				if c > 0 {
					ans++
					fmt.Println(i, s, a, b, c)
					fmt.Println(S[a+1:], S[b+1], S[c+1])
				}
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
