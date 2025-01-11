package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		max := nums[0]
		if nums[1] > max {
			max = nums[1]
		}
		return max
	}
	nums[2] = nums[2] + nums[0]

	for i := 3; i < len(nums); i++ {
		max := nums[i-2]
		if nums[i-3] > max {
			max = nums[i-3]
		}
		nums[i] = nums[i] + max
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
