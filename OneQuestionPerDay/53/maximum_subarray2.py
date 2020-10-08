from typing import List

from collections import namedtuple

status = namedtuple("Status", ["l_sum", "r_sum", "i_sum", "m_sum"])


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        return self.get(nums, 0, len(nums) - 1).m_sum

    def push_up(self, l: status, r: status) -> status:
        return status(
            max(l.l_sum, l.i_sum + r.l_sum),
            max(r.r_sum, l.r_sum + r.i_sum),
            l.i_sum + r.i_sum,
            max(l.m_sum, l.r_sum + r.l_sum, r.m_sum),
        )

    def get(self, arr: List[int], l: int, r: int) -> status:
        if l == r:
            return status(arr[l], arr[l], arr[l], arr[l])

        m = (l + r) // 2
        l_sub = self.get(arr, l, m)
        r_sub = self.get(arr, m + 1, r)
        return self.push_up(l_sub, r_sub)
