package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Atoi(n string) int {
	res := 0

	sign := 1

	nums := []rune(n)

	if nums[0] == '-' {
		sign = -1
		nums[0] = '0'
	}

	if nums[0] == '+' {
		nums[0] = '0'
	}

	for _, ch := range nums {
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

func IsPowerOf2(n int) bool {
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}

	return true
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	}

	num := Atoi(os.Args[1])

	res := IsPowerOf2(num)

	if res {
		PrintStr("true")
	} else {
		PrintStr("false")
	}

}
