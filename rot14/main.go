package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	res := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			res += string((ch-'a'+14)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			res += string((ch-'A'+14)%26 + 'A')
		} else {
			res += string(ch)
		}
	}

	return res
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
