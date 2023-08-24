package main

import "fmt"

func StrLen(str string) int {
	return len(str)
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
