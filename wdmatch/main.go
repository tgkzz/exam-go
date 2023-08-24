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

	args := os.Args[1:]

	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	j := 0

	res := ""

	for i := 0; i < len(output); i++ {
		if j < len(input) && output[i] == input[j] {
			res += string(output[i])
			j++
		}
	}

	if input == res {
		PrintStr(res)
	} else {
		z01.PrintRune('\n')
		return
	}
}
