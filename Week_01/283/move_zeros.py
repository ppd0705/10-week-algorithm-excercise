from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        count = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                count += 1
            elif count > 0:
                nums[i - count], nums[i] = nums[i], 0


