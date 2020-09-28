from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        counter = {}
        half_len = len(nums) // 2
        for n in nums:
            if n in nums:
                counter[n] += 1
            else:
                counter[n] = 1
            if counter[n] > half_len:
                return n
        return -1
