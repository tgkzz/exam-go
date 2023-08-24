package main

import (
	"fmt"
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

	if len(args) != 3 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]

	letter1 := os.Args[2]

	if len(letter1) != 1 {
		fmt.Println("Error")
	}

	letter2 := os.Args[3]

	if len(letter2) != 1 {
		fmt.Println("Error")
	}

	res := ""

	for _, ch := range input {
		if string(ch) == letter1 {
			res += letter2
		} else {
			res += string(ch)
		}
	}

	PrintStr(res)

}
