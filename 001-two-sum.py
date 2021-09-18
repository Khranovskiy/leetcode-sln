#!/bin/python3

from typing import List
import math
import os
import random
import re
import sys

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        result = []
        otherPart = dict()
        for i, num in enumerate(nums):
            second = target - num

            if second in otherPart:
                result.append(otherPart[second])
                result.append(i)
                break
            else:
                otherPart[num]=i
         

        return result
        
