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
		z01.PrintRune('\n')
		return
	}

	arg := os.Args[1]

	res := ""

	for _, ch := range arg {
		if ch >= 'a' && ch <= 'z' {
			res += string('z' - ch + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			res += string('Z' - ch + 'A')
		} else {
			res += string(ch)
		}
	}

	printStr(res)
}
