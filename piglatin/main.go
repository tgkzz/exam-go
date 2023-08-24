package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsVowel(s string) bool {
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U" {
		return true
	}

	return false
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

	res := ""

	if IsVowel(string(arg[0])) {
		res += string(arg) + "ay"
		PrintStr(res)
		return
	}

	consonant := ""
	haveVowel := false
	var end int

	for i := 0; i < len(arg); i++ {
		if !IsVowel(string(arg[i])) {
			consonant += string(arg[i])
		} else {
			end = i
			haveVowel = true
			break
		}
	}

	if !haveVowel {
		PrintStr("No vowels")
	} else {
		for i := end; i < len(arg); i++ {
			res += string(arg[i])
		}
		res += consonant + "ay"

		PrintStr(res)
	}

}
