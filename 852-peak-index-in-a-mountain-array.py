class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        low, high = 1, len(arr) - 2
        
        while low <= high:
            median = (low + high) >> 1
            
            if arr[median - 1] < arr[median] and arr[median + 1] < arr[median]:
                return median
            
            if arr[median - 1] > arr[median]:
                high = median - 1
            elif arr[median + 1] > arr[median]:
                low = median + 1

        return -1
