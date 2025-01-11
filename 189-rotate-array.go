package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(n []int, k int) {
	length := len(n)
	if length < 2 {
		return
	}
	count := k % length

	revInpl(n, 0, length-1)

	revInpl(n, 0, count-1)
	revInpl(n, count, length-1)
}

func revInpl(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
