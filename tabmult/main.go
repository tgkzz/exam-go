package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	sign := 1

	res := 0

	runes := []rune(s)

	if runes[0] == '-' {
		sign = -1
		runes[0] = '0'
	} else if runes[0] == '+' {
		runes[0] = '0'
	}

	for _, ch := range runes {
		res *= 10

		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			res += digit
		} else {
			return 0
		}
	}

	return res * sign
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

	arg := os.Args[1]

	num := Atoi(arg)

	//task: do it using printStr function
	for i := 1; i <= 9; i++ {
		fmt.Print(i)
		fmt.Print(" x ")
		fmt.Print(num)
		fmt.Print(" = ")
		fmt.Println(i * num)
	}

}
