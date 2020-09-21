class Solution:
    cache = {}
    def climbStairs(self, n: int) -> int:
        if n in self.cache:
            return self.cache[n]
        if n < 3:
            v = n
        else:
            v = self.climbStairs(n-2) + self.climbStairs(n-1)
        self.cache[n] = v
        return v
