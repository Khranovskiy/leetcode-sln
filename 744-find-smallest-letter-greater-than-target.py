class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters)-1
    
        while left <= right:
            median = (right - left) // 2 + left 
            
            if letters[median] > target:
                right = median - 1
            else:
                left = median + 1
                
        return letters[left % len(letters)]

