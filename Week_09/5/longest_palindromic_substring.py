class Solution:
    def longestPalindrome(self, s: str) -> str:
        dp = [[False] * len(s) for _ in range(len(s))]

        ans = ""
        for l in range(0, len(s)):
            for i in range(0, len(s) - l):
                j = i + l
                if l == 0:
                    dp[i][j] = True

                elif l == 1:
                    dp[i][j] = s[i] == s[j]
                else:
                    dp[i][j] = s[i] == s[j] and dp[i + 1][j - 1]

                if dp[i][j] and l + 1 > len(ans):
                    ans = s[i:j + 1]
        return ans