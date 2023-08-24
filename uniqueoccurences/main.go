package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	m := make(map[string]int)
	res := make(map[int]bool)

	arg := os.Args[1]

	for _, ch := range arg {
		m[string(ch)]++
	}

	for _, num := range m {
		_, ok := res[num]
		if !ok {
			res[num] = true
		} else {
			printStr("false")
			return
		}
	}

	printStr("true")
}
