package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	H, W int
)

func main() {
	H = ReadInt()
	W = ReadInt()

	S := make([][]byte, H)

	for i := 0; i < H; i++ {
		S[i] = ReadByteSlice(W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '.' {
				count := 0
				for a := -1; a < 2; a++ {
					for b := -1; b < 2; b++ {
						y := i + a
						x := j + b
						if y < 0 || y >= H || x < 0 || x >= W {
							continue
						}
						if S[y][x] == '#' {
							count++
						}
					}
				}
				S[i][j] = strconv.Itoa(count)[0]
			}
		}
	}
	for i := 0; i < H; i++ {
		fmt.Println(string(S[i]))
	}
}

// snipet

// INF_BIT60 = 1 << 60
var sc = bufio.NewScanner(os.Stdin)

func input() string {
	sc.Scan()
	return sc.Text()
}

func init() { sc.Buffer(make([]byte, 256), 1e9); sc.Split(bufio.ScanWords) }

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

func ReadByte() []byte {
	v := []byte(Read())
	return v
}

func ReadByteSlice(n int) []byte {
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = ReadByte()
	}
	return b
}
