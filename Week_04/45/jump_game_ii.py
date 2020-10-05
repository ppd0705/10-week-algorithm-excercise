from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        steps = end = max_pos = 0
        for i in range(0, len(nums) - 1):
            max_pos = max(max_pos, i + nums[i])
            if i == end:
                steps += 1
                end = max_pos
        return steps
