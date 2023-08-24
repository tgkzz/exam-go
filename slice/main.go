package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	var start, end int

	switch len(nbrs) {
	case 0:
		return a
	case 1:
		start = nbrs[0]
		end = len(a)
	case 2:
		start = nbrs[0]
		end = nbrs[1]
	default:
		return nil
	}

	if start < 0 {
		start += len(a)
	}

	if end < 0 {
		end += len(a)
	}

	if start < 0 {
		start = 0
	}

	if end > len(a) {
		end = len(a)
	}

	if start > end {
		return nil
	}

	return a[start:end]
}
func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))

	fmt.Println(len(a))
}
