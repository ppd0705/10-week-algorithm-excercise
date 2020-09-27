class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        elif n == 1:
            return x
        elif n < 0:
            n = -n
            x = 1 / x

        sub = self.myPow(x, n // 2)
        if n % 2 == 0:
            return sub * sub
        else:
            return sub * sub * x
