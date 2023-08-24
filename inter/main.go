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
	if len(os.Args[1:]) != 2 {
		z01.PrintRune('\n')
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	res := ""

	m := make(map[string]bool)

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(arg2); j++ {
			_, ok := m[string(arg1[i])]
			if arg1[i] == arg2[j] && !ok {
				res += string(arg1[i])
				m[string(arg1[i])] = true
				break
			}
		}
	}

	PrintStr(res)
}
