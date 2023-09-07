package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SplitWhiteSpace(s string) []string {
	someres := []string{}

	wordCollector := ""

	for _, ch := range s {
		if ch == ' ' {
			someres = append(someres, wordCollector)
			wordCollector = ""
		} else {
			wordCollector += string(ch)
		}
	}

	if wordCollector != "" {
		someres = append(someres, wordCollector)
	}

	res := []string{}

	for _, word := range someres {
		if word != "" {
			res = append(res, word)
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
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
	}

	arg := os.Args[1]

	splitted := SplitWhiteSpace(arg)

	if len(splitted) == 1 {
		PrintStr(splitted[0])
		z01.PrintRune('\n')
		return
	}

	if len(splitted) == 2 {
		PrintStr(splitted[1])
		z01.PrintRune(' ')
		PrintStr(splitted[0])
		return
	}

	for i := 1; i < len(splitted); i++ {
		PrintStr(splitted[i])
		z01.PrintRune(' ')
	}

	PrintStr(splitted[0])

}
