package main

import (
	"fmt"
)

func main() {
	var a string
	var m rune
	fmt.Scan(&a)
	rs := []rune(a)
	for i := range rs {
		if rs[i]-48 > m {
			m = rs[i] - 48
		}
	}
	fmt.Print(m)
}
