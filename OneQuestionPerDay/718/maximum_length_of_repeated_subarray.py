from typing import List


class Solution:
    def findLength(self, A: List[int], B: List[int]) -> int:
        m, n = len(A), len(B)
        if m == 0 or n == 0:
            return 0

        dp = [[0] * (n + 1) for _ in range(m + 1)]
        res = 0
        for i in range(m):
            for j in range(n):
                if A[i] == B[j]:
                    dp[i + 1][j + 1] = dp[i][j] + 1
                else:
                    dp[i + 1][j + 1] = 0
        return res

    def test(self):
        a = [1, 2, 3, 2, 1]
        b = [3, 2, 1, 4, 7]
        ans = self.findLength(a, b)
        print(ans)


if __name__ == "__main__":
    Solution().test()
