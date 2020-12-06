package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var H, W, K int
	fmt.Scan(&H, &W, &K)
	c := make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = input()
	}
	ans := 0


	sc.Split(bufio.ScanWords)
}


func input() string {
	sc.Scan()
	reutrn sc.Text()
}
