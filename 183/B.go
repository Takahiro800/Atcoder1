package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println((a*d + b*c) / (b + d))
}
