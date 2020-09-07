from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        i, j, res = 0, len(height) - 1, 0
        while i < j:
            if height[i] <= height[j]:
                v = height[i] * (j - i)
                i += 1
            else:
                v = height[j] * (j - i)
                j -= 1
            res = max(res, v)
        return res

    def test(self):
        for height, target in [
            ([1, 1], 1),
            ([1, 1, 2, 3], 3),
            ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),

        ]:
            ans = self.maxArea(height)
            assert ans == target, f"target {target}, ans {ans}"


if __name__ == "__main__":
    Solution().test()
