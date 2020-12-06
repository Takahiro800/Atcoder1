package main

import (
	"fmt"
)

func main() {
	var N string
	fmt.Scan(&N)

	var a, b, c, ans int
	a, b, c, ans = 0, 0, 0, 0
	tmp := 0
	for _, x := range N {
		k := int(x-'0') % 3
		tmp += k
		switch k {
		case 0:
			a++
		case 1:
			b++
		case 2:
			c++
		}
	}
	mod := tmp % 3

	if mod == 0 {
		ans = 0
	} else if mod == 1 {
		if b > 0 {
			ans = 1
		} else if c > 1 {
			ans = 2
		}
	} else if mod == 2 {
		if c > 0 {
			ans = 1
		} else if b > 1 {
			ans = 2
		}
	} else {
		ans = -1
	}
	if ans > 0 && ans == len(N) {
		ans = -1
	}

	fmt.Println(ans)

}
