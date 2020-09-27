class Solution:
    def getKthMagicNumber(self, k: int) -> int:
        dp = [1] * k
        p3 = p5 = p7 = 0

        for i in range(1, k):
            f3 = 3 * dp[p3]
            f5 = 5 * dp[p5]
            f7 = 7 * dp[p7]
            dp[i] = min(f3, f5, f7)
            if f3 == dp[i]:
                p3 += 1
            if f5 == dp[i]:
                p5 += 1
            if f7 == dp[i]:
                p7 += 1
        return dp[k - 1]
