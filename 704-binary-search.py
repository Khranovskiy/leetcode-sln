class Solution:
    def searchReccursive(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1
        
        middle = len(nums) // 2
        
        if nums[middle] == target:
            return middle
        if target < nums[middle]:
            return self.search(nums[:middle], target) # middle excluded
        
        index = self.search(nums[middle + 1:], target)
        if index == -1:
            return index
        return middle + 1 + index

    def search(self, nums: List[int], target: int) -> int:
        low, high = 0, len(nums)-1
    
        while low <= high:
            # median = (low + high) >>> 1
            median = (high - low) // 2 + low 
            
            if nums[median] < target:
                low = median + 1
            elif nums[median] == target:
                return median
            else:
                high = median - 1
        return -1
            
