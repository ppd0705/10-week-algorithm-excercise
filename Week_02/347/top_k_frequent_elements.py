from typing import List
from collections import defaultdict
import heapq


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = defaultdict(int)
        for n in nums:
            counter[n] += 1
        heap = []
        i = 0
        for key, v in counter.items():
            if i < k:
                heapq.heappush(heap, [v, key])
                i += 1
            elif i >= k and v > heap[0][0]:
                heapq.heapreplace(heap, [v, key])
        return [h[1] for h in heap]

    def test(self):
        for nums, k, target in [
            ([1], 1, [1]),
            ([1, 1, 1, 2, 3, 3, 4], 2, [1, 3]),
        ]:
            ans = self.topKFrequent(nums, k)
            assert sorted(ans) == target, f"target: {target}, ans: {ans}"

        print("well done")


if __name__ == "__main__":
    Solution().test()
