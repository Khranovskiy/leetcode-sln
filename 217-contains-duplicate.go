package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	set := map[int]bool{}

	for _, num := range nums {
		if set[num] {
			return true
		}
		set[num] = true
	}
	return false
}

func main() {
	b := []int{1, 2, 3, 4, 1}
	fmt.Println(containsDuplicate(b))
}

