package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)

	fmt.Println(w[0] == w[2])
}
