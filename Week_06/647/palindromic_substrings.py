class Solution:
    def countSubstrings(self, s: str) -> int:
        n = len(s)
        if n < 2:
            return n
        dp = [[False] * n for _ in range(n)]
        res = 0

        for i in range(n):
            for j in range(0, i + 1):
                if s[i] == s[j] and (i - j < 2 or dp[i - 1][j + 1]):
                    dp[i][j] = True
                    res += 1
        return res
