package main

import "fmt"

func IsPowerOf2(n uint) bool {
	res := true

	if n == 1 {
		return true
	}

	for n != 1 {
		if n%2 != 0 {
			res = false
			break
		}

		n /= 2
	}

	return res
}

func main() {
	fmt.Println(IsPowerOf2(2)) // true
	fmt.Println(IsPowerOf2(8)) // true

	fmt.Println(IsPowerOf2(14)) // false

	fmt.Println(IsPowerOf2(512)) // true

	fmt.Println(IsPowerOf2(7)) // false

	fmt.Println(IsPowerOf2(514)) // false

}
