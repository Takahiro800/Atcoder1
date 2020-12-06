package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ReadInt()
	S := ReadString()
	ans := 0

	for i := 0; i < N-2; i++ {
		if S[i:i+3] == "ABC" {
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
