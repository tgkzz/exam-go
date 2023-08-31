package main

import "os"

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

func SplitWhiteSpace(s string) []string {
	res := []string{}
	wordCollector := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			wordCollector += string(s[i])
		} else if wordCollector != "" {
			res = append(res, wordCollector)
			wordCollector = ""
		}
	}

	if wordCollector != "" {
		res = append(res, wordCollector)
	}

	return res
}

func ToCapLast(s string) string {
	splitted := SplitWhiteSpace(s)
	tmpres := []string{}
	res := ""

	for _, word := range splitted {
		tmpres = 
	}

}

func reversestr(s string) string {
	s = ToLower(s)
	s = ToCapLast(s)
}

func main() {
	res := []string{}

	args := os.Args[1:]

	for i, arg := range args {
		res[i] = reversestr(arg)
	}
}
