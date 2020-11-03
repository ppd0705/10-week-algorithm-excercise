class Solution:
    def hammingWeight(self, n: int) -> int:
        res = 0
        while n > 0:
            n, m = divmod(n, 2)
            res += m
        return res
