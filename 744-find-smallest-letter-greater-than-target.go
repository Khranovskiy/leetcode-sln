func nextGreatestLetter(nums []byte, target byte) byte {
   low, high := 0, len(nums) - 1

    for low <= high {
        median := (low + high) >> 1
        middle := nums[median]
    
        if middle <= target {
            low = median + 1
        } else { 
            high = median - 1
        }
    }
    return nums[low % len(nums)]
}
