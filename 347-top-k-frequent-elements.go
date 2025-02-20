package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freqM := make(map[int]int)
	for _, num := range nums {
		freqM[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqM {
		buckets[freq] = append(buckets[freq], num)
	}
	result := []int{}
	for i := len(nums); i > 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}
	return result[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [1 4]
}
