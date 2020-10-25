from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        pre = cur = 0
        for n in nums:
            cur, pre = max(cur, pre + n), cur
        return cur
