package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune('\n')
		return
	}

	res := ""

	for _, ch := range arg[0] {
		if ch >= 'A' && ch <= 'Z' {
			res += string((ch-'A'+13)%26 + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			res += string((ch-'a'+13)%26 + 'a')
		} else {
			res += string(ch)
		}
	}

	PrintStr(res)
}
