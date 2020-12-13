package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	S string
)

func main() {
	S = Read()
	if strings.HasPrefix(S, "A") && strings.LastIndex(S, "A") == 0 {
		c := strings.Index(S, "C")
		if c == strings.LastIndex(S, "C") {
			if 1 < c && c < len(S)-1 {
				if checkUpperCount(S) < 3 {
					fmt.Println("AC")
					return
				}
			}
		}

	}
	fmt.Println("WA")
}
func checkUpperCount(s string) int {
	count := 0
	runeS := []rune(s)
	for i := 0; i < len(s); i++ {
		if runeS[i] < 97 {
			count++
		}
	}
	return count
}

func pickCha(S string, i int) string {
	return string([]rune(S)[i])
}

func pickString(S string, i, j int) string {
	return string([]rune(S)[i : j+1])
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

func ReadFloatSlice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat()
	}
	return b
}

func ReadFloat() float64 {
	v, e := strconv.ParseFloat(Read(), 64)
	if e != nil {
		panic(e.Error())
	}
	return v
}

func judgeInt(x float64) bool {
	if math.Floor(x) == x {
		return true
	}
	return false
}
