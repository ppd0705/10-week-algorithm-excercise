class Solution:
    def addDigits(self, num: int) -> int:
        while num > 9:
            s = 0
            while num > 0:
                s += num % 10
                num //= 10
            num = s
        return num

    def test(self):
        for num, target in [
            (0, 0),
            (1, 1),
            (9, 9),
            (10, 1),
            (11, 2),
            (19, 1),
        ]:
            ans = self.addDigits(num)
            assert ans == target, f"target: {target}, ans: {ans}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
