from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        res = 0
        stack = []
        heights = [0] + heights + [0]

        for i, h in enumerate(heights):
            while len(stack) > 1 and h < heights[stack[-1]]:
                idx = stack.pop()
                res = max(res, heights[idx] * (i - stack[-1] - 1))
            stack.append(i)
        return res

    def test(self):
        for heights, target in [
            ([], 0),
            ([2, 1, 5, 6, 2, 3], 10),
        ]:
            ans = self.largestRectangleArea(heights)
            assert ans == target, f"target: {target}, ans {ans}"

        print("all done")


if __name__ == "__main__":
    Solution().test()
