package main

import (
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

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false

	if n < 0 {
		n = -n
		isNegative = true
	}

	reversed := ""

	for n > 0 {
		digit := n % 10
		reversed += string(digit + '0')
		n /= 10
	}

	res := ""

	for i := len(reversed) - 1; i >= 0; i-- {
		res += string(reversed[i])
	}

	if isNegative {
		res = "-" + res
		return res
	} else {
		return res
	}
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

	for i := 1; i <= 9; i++ {

		res := ""

		res += Itoa(i)
		res += " x "
		res += arg
		res += " = "
		res += Itoa(i * num)

		PrintStr(res)

	}

}
