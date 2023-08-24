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

func Plus(a, b int) {
	res := a + b

	if (res > a) != (b > 0) {
		z01.PrintRune('\n')
		return
	} else {
		PrintStr(Itoa(res))
	}
}

func Minus(a, b int) {
	Plus(a, -b)
}

func Divide(a, b int) {
	if b == 0 {
		PrintStr("No division by 0")
		return
	} else {
		PrintStr(Itoa(a / b))
	}
}

func Multy(a, b int) {
	res := a * b
	if a != 0 && res/a != b {
		z01.PrintRune('\n')
		return
	} else {
		PrintStr(Itoa(res))
	}
}

func Mod(a, b int) {
	if b == 0 {
		PrintStr("No modulo by 0")
		return
	} else {
		PrintStr(Itoa(a % b))
	}
}

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

func Atoi(s string) (int, bool) {
	res := 0
	sign := 1

	runes := []rune(s)

	if runes[0] == '-' {
		sign = -1
		runes[0] = '0'
	}

	if runes[0] == '+' {
		runes[0] = '0'
	}

	for _, ch := range runes {
		if ch >= '0' && ch <= '9' {
			tmp := res*10 + int(ch-'0')
			if tmp < res {
				return 0, true
			}
			res = tmp
		} else {
			return 0, true
		}
	}

	return res * sign, false
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	num1, err := Atoi(os.Args[1])
	if err {
		z01.PrintRune('\n')
		return
	}

	operator := os.Args[2]

	num2, err := Atoi(os.Args[3])
	if err {
		z01.PrintRune('\n')
		return
	}

	switch operator {
	case "+":
		Plus(num1, num2)
	case "-":
		Minus(num1, num2)
	case "/":
		Divide(num1, num2)
	case "*":
		Multy(num1, num2)
	case "%":
		Mod(num1, num2)
	}

}
