class Solution:
    def reverseBits(self, n: int) -> int:
        ans, pow = 0, 31

        while n > 0:
            ans += (1 & n) << pow
            pow -= 1
            n >>= 1
        return ans
