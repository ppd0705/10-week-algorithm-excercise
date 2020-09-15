from typing import List, Dict
from collections import defaultdict


class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        counter: Dict[int, int] = defaultdict(int)
        for n in nums1:
            counter[n] += 1

        res = []
        for n in nums2:
            if counter[n] > 0:
                res.append(n)
                counter[n] -= 1
        return res
