func search(nums []int, target int) int {
    low, high := 0, len(nums) - 1

    for low <= high {
        median := (low + high) >> 1
        // median = (high - low) / 2 + low 

        middle := nums[median]
        
        if middle == target {
            return median
        } 
        
        if middle < target {
            low = median + 1
        } else { 
            high = median - 1
        }
    }
    return -1
}
