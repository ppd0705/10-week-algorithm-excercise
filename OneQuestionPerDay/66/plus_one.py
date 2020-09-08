from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = True
        for i in range(len(digits) - 1, -1, -1):
            if digits[i] == 9:
                digits[i] = 0
            else:
                digits[i] += 1
                carry = False
                break
        if carry:
            digits.insert(0, 1)
        return digits

    def test(self):

        for digits, target in [
            ([1], [2]),
            ([0], [1]),
            ([9], [1, 0]),
            ([2, 9, 9], [3, 0, 0]),
        ]:
            ans = self.plusOne(digits)
            assert ans == target, f"target: {target}, ans {ans}"


if __name__ == "__main__":
    Solution().test()
