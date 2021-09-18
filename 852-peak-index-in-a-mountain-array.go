
func peakIndexInMountainArray(nums []int) int {
    low, high := 1, len(nums) - 2

    for low <= high {
        median := (low + high) >> 1

        if nums[median-1] < nums[median] && nums[median+1] < nums[median] {
            return median
        }
        if nums[median - 1] > nums[median] {
            high = median - 1
        } else if nums[median + 1] > nums[median] {
            low = median + 1
        }
    }
    return -1
}
