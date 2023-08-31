package main

import (
	"os"

	"github.com/01-edu/z01"
)

func FindFirstCh(s string) int {
	for i, ch := range s {
		if ch != ' ' {
			return i
		}
	}

	return 0
}

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

	str := os.Args[1]

	if len(str) == 0 {
		return
	}

	idx := FindFirstCh(str)

	res := ""
	for i := idx; i < len(str); i++ {
		if str[i] != ' ' {
			res += string(str[i])
		} else {
			break
		}
	}

	PrintStr(res)
}
