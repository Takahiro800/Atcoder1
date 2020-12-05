package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var a, b, c, d int
	a = int(s[0] - '0')
	b = int(s[1] - '0')
	c = int(s[2] - '0')
	d = int(s[3] - '0')

	for _, i := range []int{-1, 1} {
		for _, j := range []int{-1, 1} {
			for _, k := range []int{-1, 1} {
				if a+i*b+j*c+k*d == 7 {
					op := map[int]string{-1: "-", 1: "+"}
					fmt.Print(a)
					fmt.Print(op[i])
					fmt.Print(b)
					fmt.Print(op[j])
					fmt.Print(c)
					fmt.Print(op[k])
					fmt.Print(d)
					fmt.Println("=7")
					return
				}
			}
		}
	}
}
