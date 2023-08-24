package main

import "github.com/01-edu/z01"

func main() {
	asd := "98765432310"

	for _, ch := range asd {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
