package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var array [100][2]float64
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i][0], &array[i][1])
	}
	ans := 0.0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if v := math.Hypot(array[i][0]-array[j][0], array[i][1]-array[j][1]); ans < v {
				ans = v
			}
		}
	}
	fmt.Println(ans)
}
