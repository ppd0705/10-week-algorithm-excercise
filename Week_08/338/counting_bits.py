from typing import List


class Solution:
    def countBits(self, num: int) -> List[int]:
        res = [0] * (num + 1)
        for i in range(1, num + 1):
            # res[i] = (i & 1) + res[i >> 1]
            res[i] = i % 2 + res[i // 2]
        return res
