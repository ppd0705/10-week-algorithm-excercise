class Solution:
    def isPowerOfFour(self, num: int) -> bool:
        while num > 0 and num % 4 == 0:
            num //= 4
        return num == 1
