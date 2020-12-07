package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ReadInt()
	ans := 0
	for i := 1; i < N+1; i++ {
		a := strconv.Itoa(i)
		if len(a)%2 == 1 {
			ans++
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
