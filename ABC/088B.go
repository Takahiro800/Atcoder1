package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, alice, bob int
	fmt.Scan(&n)
	cards := make([]int, n)
	fmt.Println(cards)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
		fmt.Println(cards)
	}
	sort.SliceStable(cards, func(i, j int) bool {
		if cards[i] > cards[j] {
			return true
		}
		return false
	})
	for i := 0; i < n; i += 2 {
		alice += cards[i]
		if i+1 < n {
			bob += cards[i+1]
		}
	}
	fmt.Println(alice - bob)
}
