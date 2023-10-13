package main

func Divisions(n, divisor int) int {
	count := 0
	for n > 0 {
		n /= divisor
		count++
	}

	return count + 1
}
