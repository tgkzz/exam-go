package main

import "fmt"

func Atoi(s string) int {
	sign := 1
	res := 0

	runes := []rune(s)

	if runes[0] == '+' {
		runes[0] = '0'
	}

	if runes[0] == '-' {
		sign = -1
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

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
