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

func deleteSpace(s []string) []string {
	res := []string{}

	for _, word := range s {
		if len(word) != 0 || word == " " {
			res = append(res, word)
		}
	}

	return res
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune('\n')
		return
	}

	splitted := SplitWhiteSpace(arg[0])

	deleted := deleteSpace(splitted)

	if len(deleted) == 0 {
		z01.PrintRune(' ')
		z01.PrintRune('\n')
		return
	}

	printStr(deleted[len(deleted)-1])

}
