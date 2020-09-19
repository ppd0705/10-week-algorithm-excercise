class Solution:
    def addDigits(self, num: int) -> int:
        if num == 0:
            return 0
        return (num - 1) % 9 + 1

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
