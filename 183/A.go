package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(max(x, 0))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
