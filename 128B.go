package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	N int
)

type restaurant struct {
	place string
	point int
	num   int
}

func main() {
	N = ReadInt()
	restaurant_list := make([]restaurant, N)
	for i := 0; i < N; i++ {
		restaurant_list[i].place = Read()
		restaurant_list[i].point = ReadInt()
		restaurant_list[i].num = i + 1
	}
	sort.Slice(restaurant_list, func(i, j int) bool {
		if restaurant_list[i].place == restaurant_list[j].place {
			return restaurant_list[i].point > restaurant_list[j].point
		} else {
			return restaurant_list[i].place < restaurant_list[j].place
		}
	})
	for i, _ := range restaurant_list {
		fmt.Println(restaurant_list[i].num)
	}
}

// snipet
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
