package main

import "fmt"

func ItoaBase(value, base int) string {

	if value == 0 {
		return "0"
	}

	if base < 2 || base > 16 {
		return ""
	}

	isNegative := false
	if value < 0 {
		isNegative = true
		value = -1 * value
	}

	basemap := make(map[int]rune)

	i := 0
	for i < 10 {
		if i > base {
			break
		}
		basemap[i] = rune(i + '0')
		i++
	}

	for i := 10; i < 16; i++ {
		basemap[i] = rune('A' + (i - 10))
	}

	res := ""

	for value > 0 {
		digit := value % base
		res = string(basemap[digit]) + res
		value /= base
	}

	if isNegative {
		return "-" + res
	}

	return res
}

func main() {
	fmt.Println(ItoaBase(0, 2))
	fmt.Println(ItoaBase(5, 2))
	fmt.Println(ItoaBase(10, 10))
	fmt.Println(ItoaBase(255, 16))
}
