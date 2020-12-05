package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(judge(a, b))
}

func judge(a, b int) string {
	if a > 0 {
		return "Positive"
	} else if a <= 0 && b >= 0 {
		return "Zero"
	} else if (b-a)%2 == 0 {
		return "Negative"
	} else {
		return "Positive"
	}
}
