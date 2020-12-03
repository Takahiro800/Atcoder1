package main

import "fmt"

func main() {
	var K, S, ans int
	fmt.Scan(&K, &S)
	for i := 0; i <= K; i++ {
		for j := 0; j <= K; j++ {
			if l := S - i - j; l >= 0 && l <= K {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
