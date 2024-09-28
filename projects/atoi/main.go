package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	rs := []rune(a)
	for i := range rs {
		fmt.Print((rs[i] - 48) * (rs[i] - 48))
	}
}
