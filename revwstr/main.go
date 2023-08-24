package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SplitWhiteSpace(s string) []string {
	res := []string{}
	wordCollector := ""

	for _, ch := range s {
		if ch != ' ' {
			wordCollector += string(ch)
		} else {
			res = append(res, wordCollector)
			wordCollector = ""
		}
	}

	if wordCollector != "" {
		res = append(res, wordCollector)
	}

	return res
}

func printStr(s string) {
	for _, chj := range s {
		z01.PrintRune(chj)
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	arg := os.Args[1]

	if arg == "" {
		z01.PrintRune('\n')
		return
	}

	split := SplitWhiteSpace(arg)

	for i := len(split) - 1; i >= 0; i-- {
		printStr(split[i])
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
