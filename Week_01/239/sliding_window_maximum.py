from typing import List
from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq = deque()
        res = []
        for i, n in enumerate(nums):
            while dq and dq[0] <= i - k:
                dq.popleft()
            while dq and nums[dq[-1]] < n:
                dq.pop()
            dq.append(i)
            if i >= k - 1:
                res.append(nums[dq[0]])
        return res

    def test(self):
        for nums, k, target in [
            ([], 3, []),
            ([1, 2], 1, [1, 2]),
            ([1, 3, -1, -3, 5, 3, 6, 7], 3, [3, 3, 5, 5, 6, 7]),
        ]:
            ans = self.maxSlidingWindow(nums, k)
            assert ans == target, f"target: {target}, ans: {ans}"
        print("all done")


if __name__ == "__main__":
    Solution().test()
