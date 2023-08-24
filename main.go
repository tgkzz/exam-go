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

	input := arg[0]

	res := ""

	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			res += string('z' - ch + 'a')
		} else {
			res += string(ch)
		}
	}

	PrintStr(res)

}
