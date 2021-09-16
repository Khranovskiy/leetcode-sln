type NumArray struct {
    sums []int
}

func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums), len(nums))
    
    current_sum := 0
    for i, num := range nums {
        current_sum += num
        sums[i] = current_sum
    }
    // alternative:
    // sums[0] = nums[0]
    // for i := 1; i < len(nums); i++ {
    //     sums[i] = sums[i-1] + nums[i]
    // }
    return NumArray{ sums }
}

func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.sums[right]
    }
    return this.sums[right] - this.sums[left - 1]
}