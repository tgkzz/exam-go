package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func Atoi(s string) (int, bool) {
	runes := []rune(s)
	sign := 1
	res := 0

	if runes[0] == '+' {
		runes[0] = 0
	}
	if runes[0] == '-' {
		sign = -1
		runes[0] = 0
	}

	for _, ch := range s {
		res *= 10
		if ch >= '0' && ch <= '9' {
			res += int(ch - '0')
		} else {
			return 0, false
		}
	}

	return res * sign, true
}

func ToBin(n int) string {
	reversed := ""

	for n > 0 {
		digit := n % 2
		reversed += string(digit + '0')
		n /= 2
	}

	res := ""
	for i := len(reversed) - 1; i >= 0; i-- {
		res += string(reversed[i])
	}

	return res
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	arg := os.Args[1]

	num, ok := Atoi(arg)
	if !ok {
		PrintStr("00000000")
	}

	res := ToBin(num)

	if len(res) != 8 {
		for i := len(res); i <= 8; i++ {
			res = "0" + res
		}
		PrintStr(res)
	} else {
		PrintStr(res)
	}
}
