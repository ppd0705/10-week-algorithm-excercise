class Solution:
    def longestValidParentheses(self, s: str) -> int:
        dp = [0] * len(s)
        ans = 0
        for i in range(1, len(s)):
            if s[i] == ')':
                if s[i - 1] == '(':
                    if i > 1:
                        dp[i] = dp[i - 2] + 2
                    else:
                        dp[i] = 2
                elif dp[i - 1] > 0 and i - dp[i - 1] - 1 >= 0 and s[i - dp[i - 1] - 1] == '(':
                    if i - dp[i - 1] - 2 > 0:
                        dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2
                    else:
                        dp[i] = dp[i - 1] + 2
                ans = max(ans, dp[i])
        return ans

    def test(self):
        s = "(())"
        ans = self.longestValidParentheses(s)
        target = 4
        assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
