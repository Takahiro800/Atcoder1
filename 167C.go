package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	n := sc.readInt()
	m := sc.readInt()
	x := sc.readInt()

	arry := make([][]int, n)
	for i := 0; i < n; i++ {
		cs := getNums(sc, m+1)
		arry[i] = cs
	}
	var minnum = -1
	for bits := 0; bits < (1 << uint64(n)); bits++ {
		var sum = make([]int, m+1)
		// bitsの各要素についてのループ
		for i := 0; i < n; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bits>>uint64(i))&1 == 1 {
				for j := 0; j < m+1; j++ {
					sum[j] += arry[i][j]
				}
				// 何かしらの処理
			}
			var flg bool = true
			for j := 1; j < m+1; j++ {
				if sum[j] >= x {
					continue
				}
				flg = false
			}
			if flg {
				if minnum == -1 || minnum > sum[0] {
					minnum = sum[0]
				}
			}
		}
	}
	fmt.Println(minnum)
}

func (sc *scanner) readInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func getNums(sc *scanner, len int) (nums []int) {
	var a = make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt()
	}
	return a
}

func newScanner() *scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &scanner{bufSc}
}

type scanner struct {
	bufScanner *bufio.Scanner
}
