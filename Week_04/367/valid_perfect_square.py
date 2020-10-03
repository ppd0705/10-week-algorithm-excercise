class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        l, r = 0, num
        while l < r:
            m = (l + r) // 2
            square = m * m
            if square == num:
                return True
            elif square < num:
                l = m + 1
            else:
                r = m - 1

        return False

    def test(self):
        assert self.isPerfectSquare(16)
        print("well done")


if __name__ == "__main__":
    Solution().test()
