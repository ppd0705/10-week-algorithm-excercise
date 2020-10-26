class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == '0':
            return 0
        n = len(s)
        dp = [0] * (n + 1)
        dp[0] = 1
        dp[1] = 1

        for i in range(1, n):
            if s[i] == '0':
                if s[i - 1] in ('1', '2'):
                    dp[i + 1] = dp[i - 1]
                else:
                    return 0
            else:
                if s[i - 1] == '1' or (s[i - 1] == '2' and '1' <= s[i] <= '6'):
                    dp[i + 1] = dp[i] + dp[i - 1]
                else:
                    dp[i + 1] = dp[i]
        return dp[-1]

    def test(self):
        for s, target in [
            ("0", 0),
            ("100", 0),
            ("27", 1),
            ("22", 2),
            ("123", 3),
            ("1221", 5),
        ]:
            ans = self.numDecodings(s)
            assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
