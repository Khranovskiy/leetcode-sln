package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

import (
	"fmt"
)

func main() {
	//	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	//fmt.Println(removeDuplicates([]int{0,0}))
	fmt.Println(removeDuplicates([]int{1, 2}))
}

func removeDuplicates(nums []int) int {
	// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
	// asc order
	// inplace remove duplicates
	// ret number of unique elements in nums

	//  0   1 2 3 4 5 6
	// -100 0 2 2 2 5 5
	//        k     r
	// -100 0 2 5 2 2 5  swap
	//          k     r
	// left 0
	// right 1
	// if v[r] > v[l] {
	//   if r > l + 1
	//       swap v[l+1] v[r]
	//  l++
	//  r++
	// }

	// if val[right] == val[left]
	//      inc(right)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	k := 0
	r := 1
	for r < len(nums) {
		if nums[r] > nums[k] {
			if r > k+1 {
				swap(nums, k+1, r)
			}
			k++
			r++
		} else if nums[r] == nums[k] {
			r++
		}
	}
	fmt.Println(nums[:k+1])
	return k + 1
}

func swap(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}

// 20 minutes to code
