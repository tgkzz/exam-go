package main

import "fmt"

func IsPrime(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	for nb > 2 {
		if IsPrime(nb) {
			return nb
		}

		nb--
	}

	return 2
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(96))
}
