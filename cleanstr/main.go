package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SplitWhiteSpace(s string) []string {
	res := []string{}
	wordCollector := ""

	for _, ch := range s {
		if ch != '\n' && ch != ' ' && ch != '\t' {
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
}

func main() {
	if len(os.Args[1:]) != 1 || len(os.Args[1]) == 0 {
		z01.PrintRune('\n')
		return
	}

	str := os.Args[1]

	splitted := SplitWhiteSpace(str)

	for i, word := range splitted {
		if i != len(splitted)-1 {
			PrintStr(word + " ")
		} else {
			PrintStr(word)
		}
	}

}
