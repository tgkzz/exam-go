package main

import "github.com/01-edu/z01"

func main() {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))
				if i != 2 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
