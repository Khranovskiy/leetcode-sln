func twoSum(nums []int, target int) []int {
    other := make(map[int]int)
    
    for index, num := range nums {
        diff := target - num
        
        if otherIndex, ok := other[diff]; ok {
            return []int{otherIndex, index}
        }
        other[num] = index
    }
    return []int{}
}
