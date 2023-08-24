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
	if len(os.Args[1:]) != 1 {
		return
	}

	arg := os.Args[1]

	res := ""

	for _, ch := range arg {
		if ch >= 'a' && ch <= 'z' {
			for i := 0; i < int(ch-'a'+1); i++ {
				res += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			for i := 0; i < int(ch-'A'+1); i++ {
				res += string(ch)
			}
		} else {
			res += string(ch)
		}
	}

	PrintStr(res)

}
