package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
func maxSubArray(nums []int) int {
	max_sum := nums[0]
	cur_sum := nums[0]

	for _, num := range nums[1:] {
		cur_sum = max(cur_sum+num, num)
		max_sum = max(cur_sum, max_sum)
	}
	return max_sum
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
