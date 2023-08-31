package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
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

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	arg := os.Args[1]

	num, ok := Atoi(arg)
	if !ok {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if num <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for num >= 2 {
		if IsPrime(num) {
			sum += num
		}
		num--
	}

	res := Itoa(sum)

	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
