package main

import "fmt"

func main() {
	sur := 9
	time := 0
	flag := 0
	array := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
		sub := array[i] % 10
		if sub != 0 && sub <= sur {
			sur = sub
			flag = 1
		}
		if sub == 0 {
			time += array[i]
		} else {
			time += array[i] + (10 - sub)
		}
	}
	if flag == 1 {
		time = time - (10 - sur)
	}
	fmt.Println(time)
}
