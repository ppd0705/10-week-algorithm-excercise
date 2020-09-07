class Solution:
    def climbStairs(self, n: int) -> int:
        if n < 3:
            return n
        f1, f2 = 1, 2
        for i in range(3, n+1):
            f2, f1, = f1 + f2, f2
        return f2

    def test(self):
        for n, target in [
            (1, 1),
            (2, 2),
            (3, 3),
            (4, 5),
            (5, 8),
        ]:
            ans = self.climbStairs(n)
            assert ans == target, f"n {n}, ans {ans}, target {target}"


if __name__ == "__main__":
    Solution().test()
