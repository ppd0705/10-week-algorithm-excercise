from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        tail = []
        for i in range(len(nums)):
            if i == 0 or nums[i] > tail[-1]:
                tail.append(nums[i])
            else:
                l, r = 0, len(tail)
                while l < r:
                    m = (l + r) // 2
                    if tail[m] < nums[i]:
                        l = m + 1
                    else:
                        r = m
                tail[l] = nums[i]
        return len(tail)
