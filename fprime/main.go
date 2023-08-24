package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var isNegative bool

	digits := []rune{}

	res := ""

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit+'0'))
		n /= 10
	}

	if isNegative {
		digits = append(digits, '-')
	}

	for i := len(digits) - 1; i >= 0; i-- {
		res += string(digits[i])
	}

	return res
}

func PosFprime(n int) int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return i
		}
	}

	return n
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	res := 0

	sign := 1

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

	return sign * res
}

func printStr(s string) {
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

	res := ""

	for num > 0 {
		factor := PosFprime(num)
		res += Itoa(factor) + "*"
		num /= factor
	}

	if len(res) > 0 {
		res = res[:len(res)-1]
	}

	printStr(res)

}
