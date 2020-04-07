package main

import (
	"fmt"
	"strconv"
	"strings"
)

//https://github.com/austingebauer/go-leetcode/tree/master/move_zeroes_283
func moveZeroes(nums []int) {
	placement := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			hold := nums[i]
			nums[i] = nums[placement]
			nums[placement] = hold
			placement++
		}
	}
}

func main() {
	values := []int{0, 1, 0, 3, 12}
	moveZeroes(values)
	valuesText := []string{}

	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, ", ")
	fmt.Printf("[%s]\n", result)
}
