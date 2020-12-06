package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	ans := 201

	t := []int{}
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		t = append(t, a)
	}
	//bitsがn個の要素の各パターンを表す
	for bits := 0; bits < (1 << uint64(N)); bits++ {
		// bitsの各要素についてのループ
		sum_a := 0
		sum_b := 0
		for i := 0; i < N; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bits>>uint64(i))&1 == 1 {
				// 何かしらの処理
				sum_a += t[i]
			} else {
				sum_b += t[i]
			}
		}
		ans = min(ans, max(sum_a, sum_b))
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
