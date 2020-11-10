class Solution:
    def longestValidParentheses(self, s: str) -> int:
        ans = 0
        left = right = 0

        for i in range(len(s)):
            if s[i] == '(':
                left += 1
            else:
                right += 1
            if left == right:
                ans = max(ans, 2 * left)
            elif left < right:
                left = right = 0
        left = right = 0

        for i in range(len(s) - 1, -1, -1):
            if s[i] == '(':
                left += 1
            else:
                right += 1

            if left == right:
                ans = max(ans, 2 * left)
            elif right < left:
                left = right = 0
        return ans

    def test(self):
        s = "(())"
        ans = self.longestValidParentheses(s)
        target = 4
        assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
