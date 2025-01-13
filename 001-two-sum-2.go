func twoSum(nums []int, target int) []int {
	// return indices of the two numbers such that they add up to target.

	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if secondIndex, ok := indexMap[complement]; ok {
			return []int{i, secondIndex}
		}
		indexMap[num] = i
	}
	return nil
}

// 9 min