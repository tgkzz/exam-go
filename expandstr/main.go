package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Divided(s string) []string {
	res := []string{}
	wordCollector := ""

	for _, ch := range s {
		if ch != ' ' {
			wordCollector += string(ch)
		} else {
			if wordCollector != "" {
				res = append(res, wordCollector)
				wordCollector = ""
			}
		}
	}

	if wordCollector != "" {
		res = append(res, wordCollector)
	}

	return res
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	arg := os.Args[1]

	divideSpaces := Divided(arg)

	res := ""
	for i, word := range divideSpaces {
		res += word
		if i != len(divideSpaces)-1 {
			res += "   "
		}
	}

	PrintStr(res)
}
