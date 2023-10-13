package main

import "fmt"

func Split(s, sep string) []string {
	res := []string{}

	wordCollector := ""

	sepLen := len(sep)

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			res = append(res, wordCollector)
			wordCollector = ""
			i += sepLen - 1
		} else {
			wordCollector += string(s[i])
		}
	}
	res = append(res, wordCollector)

	return res
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
