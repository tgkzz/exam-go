package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printNum(a, b int) string {
	if a < b {
		a, b = b, a
	}

	res := ""

	for i := a; i >= b; i-- {
		tmp := strconv.Itoa(i)
		res += tmp
		if i != b {
			res += " "
		}
	}

	return res
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	str := printNum(num1, num2)

	PrintStr(str)
}
