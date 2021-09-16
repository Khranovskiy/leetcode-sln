class Solution:

    # complexity:
    # space - O(1)
    # runtime - O(n)
    # modification input data = yes
def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
    cursor = 0
    n = len(nums)
    while cursor < n:
        value = nums[cursor]
        position = value - 1

        if nums[cursor] != nums[position] : 
            nums[cursor], nums[position] = nums[position], nums[cursor]
        else:
            cursor += 1

    result = []
    for i in range(n):
        val = i + 1
        if val != nums[i]:
            result.append(val)

    return result

    # complexity:
    # space - O(n)
    # runtime - O(n)
    # modification input data = no
    def findDisappearedNumbers2(self, nums: List[int]) -> List[int]:
        setNums = set(nums)
        result = []
        for i in range(1, len(nums) + 1): # [1..n]
            if i not in setNums:
                result.append(i)

        return result


# tests
# t([1,1]) => [2]
# t([1]) => []
# t([4,3,2,7,8,2,3,1]) => [5,6]
# t([]) => [] -- invalid test case by constratints


