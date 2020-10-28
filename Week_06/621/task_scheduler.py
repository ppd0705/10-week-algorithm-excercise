from typing import List


class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        m = len(tasks)
        if m < 2 or n == 0:
            return m
        counter = [0] * 26

        for t in tasks:
            counter[ord(t) - ord('A')] += 1

        counter.sort()

        max_count = counter[-1]

        res = (max_count - 1) * (n + 1) + 1
        for c in counter[:-1]:
            if c == max_count:
                res += 1
        return max(res, m)
