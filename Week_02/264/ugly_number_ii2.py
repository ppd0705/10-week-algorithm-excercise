class Solution:
    def nthUglyNumber(self, n: int) -> int:
        dp = [0] * n
        dp[0] = 1

        two = three = five = 0

        for i in range(1, n):
            num = min(2 * dp[two], 3 * dp[three], 5 * dp[five])
            dp[i] = num
            if 2 * dp[two] == num:
                two += 1
            if 3 * dp[three] == num:
                three += 1
            if 5 * dp[five] == num:
                five += 1
        return dp[-1]

    def test(self):
        for n, target in [
            (1, 1),
            (10, 12),
        ]:
            ans = self.nthUglyNumber(n)
            assert ans == target, f"n: {n}, target: {target}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
