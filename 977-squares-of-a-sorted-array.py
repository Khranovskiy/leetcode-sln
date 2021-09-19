def sortedSquares(self, nums: List[int]) -> List[int]:
    result = [0] * len(nums)
    low, high = 0, len(nums) - 1
    ri = high
    bigger = 0

    while ri >= 0:
        leftVal = nums[low] ** 2
        rightVal = nums[high] ** 2

        if leftVal > rightVal:
            bigger = leftVal
            low += 1
        elif rightVal >= leftVal:
            bigger = rightVal
            high -= 1

        result[ri] = bigger
        ri -= 1
    return result
