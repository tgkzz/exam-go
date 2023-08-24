package main

import (
	"fmt"
)

func Capitalize(s string) string {
	capNext := true
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if capNext {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 'a' - 'A'
			}
			capNext = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 'a' - 'A'
		}
		if !(runes[i] >= 'a' && runes[i] <= 'z') && !(runes[i] >= 'A' && runes[i] <= 'Z') && !(runes[i] >= '0' && runes[i] <= '9') {
			capNext = true
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? how+are+things+4you?"))
}
