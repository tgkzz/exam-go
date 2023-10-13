package main

import "fmt"

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

func main() {
	fmt.Println(Itoa(-123))
}
