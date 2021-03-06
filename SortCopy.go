package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(left, right []int) (ret []int) {
	ret = []int{}
	for len(left) > 0 && len(right) > 0 {
		var x int
		// ソート済みのふたつのスライスからより小さいものを選んで追加していく (これがソート処理)
		if right[0] > left[0] {
			x, left = left[0], left[1:]
		} else {
			x, right = right[0], right[1:]
		}
		ret = append(ret, x)
	}
	// 片方のスライスから追加する要素がなくなったら残りは単純に連結できる (各スライスは既にソートされているため)
	ret = append(ret, left...)
	ret = append(ret, right...)
	return
}

func sort(left, right []int) (ret []int) {
	// ふたつのスライスをそれぞれ再帰的にソートする
	if len(left) > 1 {
		l, r := split(left)
		left = sort(l, r)
	}
	if len(right) > 1 {
		l, r := split(right)
		right = sort(l, r)
	}

	// ソート済みのふたつのスライスをひとつにマージする
	ret = merge(left, right)
	return
}

func split(values []int) (left, right []int) {
	// スライスを真ん中でふたつに分割する
	left = values[:len(values)/2]
	right = values[len(values)/2:]
	return
}

func Sort(values []int) (ret []int) {
	left, right := split(values)
	ret = sort(left, right)
	return
}

func main() {

	values := make([]int, N)
	// ソートして結果を出力する
	sortedValues := Sort(values)
	fmt.Println(sortedValues)
}
