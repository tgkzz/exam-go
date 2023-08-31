package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ToLower(s string) string {
	res := ""

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			res += string(ch - 'A' + 'a')
		} else {
			res += string(ch)
		}
	}

	return res
}

func ToCap(s string) string {
	res := ""

	for _, ch := range s {
		if 
	}

}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args[1:]) != 0 {
		return
	}

	args := os.Args[1:]

	for i, arg := range args {
		args[i] = ToLower(arg)
	}

	for i, arg := range args {
		args[i] = ToCap(arg)
	}
}
