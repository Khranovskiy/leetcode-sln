func twoSum(n []int, target int) []int {
	// array of integers numbers
	// sorted in non-decreasing order
	// You may not use the same element twice.

	// must use only constant extra space.

	// [2,7,11,15] target 9
	// two pointers

	// 2 + 15 > 9
	// right value greater target -> rightIndex--

	sum, left, right := 0, 0, len(n)-1

	for left < right {
		sum = n[left] + n[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 6 min