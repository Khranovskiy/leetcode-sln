package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(sortByBits([]int{8, 4, 2, 1}))
}

func sortByBits(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	slices.SortFunc(result, func(a, b int) int {
		bitsA := countBits(a)
		bitsB := countBits(b)
		// should return a negative number when a < b
		if bitsA < bitsB {
			return -1

		} else if bitsA > bitsB {
			// positive number when  a > b
			return 1
		}

		return a - b
	})
	return result
}

func countBits(n int) int {
	var res int
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
