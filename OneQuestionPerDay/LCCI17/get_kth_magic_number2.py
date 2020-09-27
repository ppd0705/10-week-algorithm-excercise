import heapq


class Solution:
    def getKthMagicNumber(self, k: int) -> int:
        h = [1]
        heapq.heapify(h)
        visited = {1: True}
        i = 0
        n = 1
        while i < k:
            i += 1
            n = heapq.heappop(h)
            for f in [3, 5, 7]:
                m = f * n
                if m not in visited:
                    heapq.heappush(h, m)
                    visited[m] = True
        return n
