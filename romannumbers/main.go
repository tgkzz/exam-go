package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IntToRoman(n int) string {
	res := ""

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	expression := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	for i := 0; i < len(val); i++ {
		for n >= val[i] {
			n -= val[i]
			if res != "" {
				res += "+"
			}
			if symbol[i] == expression[i] {
				res += symbol[i]
			} else {
				res += expression[i]
			}
		}
	}

	return res
}

func Atoi(s string) (int, bool) {
	res := 0
	sign := 1

	runes := []rune(s)

	if runes[0] == '+' {
		runes[0] = '0'
	}
	if runes[0] == '-' {
		runes[0] = '0'
		sign = -1
	}

	for _, ch := range runes {
		res *= 10
		if ch >= '0' && ch <= '9' {
			res += int(ch - '0')
		} else {
			return 0, false
		}
	}

	return res * sign, true
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

	num, ok := Atoi(arg)
	if !ok {
		PrintStr("ERROR: cannot convert to roman digit")
		return
	}

	if num <= 0 || num >= 4000 {
		PrintStr("ERROR: cannot convert to roman digit")
		return
	}

	res := IntToRoman(num)

	PrintStr(res)
}
