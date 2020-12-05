package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var N, M, k, a, ans int
	fmt.Scan(&N, &M)

	f := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&a)
			f[a]++
		}
	}
	for _, v := range f {
		if v == N {
			ans++
		}
	}
	fmt.Println(ans)
}
