package main

import "fmt"

func Max(a []int) int {
	res := -129038120

	for _, num := range a {
		if num > res {
			res = num
		}
	}

	return res
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
