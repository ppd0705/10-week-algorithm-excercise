from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:

        stack = []
        res = 0

        for i, h in enumerate(height):
            while stack and height[stack[-1]] < h:
                top = stack.pop()
                if not stack:
                    break
                left = stack[-1]
                res += (i - left - 1) * (min(h, height[left]) - height[top])
            stack.append(i)
        return res

    def test(self):
        for height, target in [
            ([], 0),
            ([1, 1], 0),
            ([1, 0, 0], 0),
            ([1, 0, 1], 1),
            ([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1], 6),
        ]:
            ans = self.trap(height)
            assert ans == target, f"target: {target}, ans {ans}"
        print("all done")


if __name__ == "__main__":
    Solution().test()
