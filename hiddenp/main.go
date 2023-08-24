package main

import (
	"fmt"
	"os"
)

func hiddenp(s1, s2 string) int {
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s1[j] == s2[i] {
			j++
		}
	}
	if j == len(s1) {
		return 1
	}
	return 0
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	fmt.Println(hiddenp(args[0], args[1]))
}
