package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	res := 0

	sign := 1

	runes := []rune(s)

	if runes[0] == '+' {
		runes[0] = '0'
	} else if runes[0] == '-' {
		runes[0] = '0'
		sign = -1
	}

	for _, ch := range runes {
		res *= 10
		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			res += digit
		} else {
			return 0, true
		}
	}

	return sign * res, false
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	num, mistake := Atoi(os.Args[1])

	if mistake {
		PrintStr("ERROR")
		return
	}

	reversed := ""
	res := ""

	hex := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}

	for num > 0 {
		digit := num % 16
		reversed += hex[digit]
		num /= 16
	}

	for i := len(reversed) - 1; i >= 0; i-- {
		res += string(reversed[i])
	}

	PrintStr(res)
}
