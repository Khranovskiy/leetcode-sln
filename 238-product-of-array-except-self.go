package main

import "fmt"

func main() {
	n := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(n))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] = r * res[j]
		r *= nums[j]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	rightProduct[len(nums)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		rightProduct[j] = rightProduct[j+1] * nums[j+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = leftProduct[i] * rightProduct[i]
	}
	return nums
}
