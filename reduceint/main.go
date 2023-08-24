package main

import "github.com/01-edu/z01"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
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

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) != 2 {
		return
	}

	PrintStr(Itoa(f(a[0], a[1])))
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
