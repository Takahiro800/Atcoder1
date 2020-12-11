package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)
	f := make(map[byte]int)

	if len(w)%2 == 1 {
		fmt.Println("No")
		return
	}

	for i := 0; i < len(w); i++ {
		if _, e := f[w[i]]; e {
			f[w[i]]++
		} else {
			f[w[i]] = 1
		}
	}
	for key, _ := range f {
		if f[key]%2 == 1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
