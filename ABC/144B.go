package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	N := ReadInt()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == N {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
