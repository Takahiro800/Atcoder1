package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func input() string {
	sc.Scan()
	return sc.Text()
}

func init()        { sc.Buffer(make([]byte, 256), 1e9); sc.Split(bufio.ScanWords) }
func Read() string { sc.Scan(); return sc.Text() }

func ReadInt() int {
	v, e := strconv.Atoi(Read())
	if e != nil {
		panic(e.Error())
	}
	return v
}

func ReadString() string {
	if !sc.Scan() {
		panic(sc.Err())
	}
	return sc.Text()
}

func main() {
	S := ReadString()
	ans := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'A' || S[i] == 'C' || S[i] == 'G' || S[i] == 'T' {
			min := 1
			for j := i + 1; j < len(S); j++ {
				if S[j] == 'A' || S[j] == 'C' || S[j] == 'G' || S[j] == 'T' {
					min++
					continue
				}
				break
			}
			if ans < min {
				ans = min
			}
		}

	}
	fmt.Println(ans)
}
