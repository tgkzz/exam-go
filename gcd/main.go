package main

import (
	"os"

	"github.com/01-edu/z01"
)

func gcd(a, b int) int {
	var min int
	if a > b {
		min = b
	} else {
		min = a
	}

	res := 0

	for i := min; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
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

func Itoa(n int) string {
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

	return res
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}

	num1, ok1 := Atoi(os.Args[1])

	if !ok1 {
		return
	}
	num2, ok2 := Atoi(os.Args[2])

	if !ok2 {
		return
	}

	resInt := gcd(num1, num2)

	resStr := Itoa(resInt)

	PrintStr(resStr)
}
