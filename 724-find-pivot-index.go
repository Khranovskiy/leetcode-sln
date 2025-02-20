package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	leftSum := 0

	for i, num := range nums {
		// Right sum is totalSum - leftSum - nums[i]
		rightSum := totalSum - leftSum - num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}

func main() {
	// Example input
	nums := []int{1, 7, 3, 6, 5, 6}

	// Find the pivot index
	result := pivotIndex(nums)
	fmt.Printf("Pivot Index: %d\n", result)
}
