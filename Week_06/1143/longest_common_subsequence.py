class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        m, n = len(text1), len(text2)
        cache = {}

        def dp(i: int, j: int) -> int:
            if i < 0 or j < 0:
                return 0
            p = (i, j)
            if p in cache:
                return cache[p]
            if text1[i] == text2[j]:
                cache[p] = dp(i - 1, j - 1) + 1
            else:
                cache[p] = max(dp(i - 1, j), dp(i, j - 1))
            return cache[p]

        return dp(m - 1, n - 1)
