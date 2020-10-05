from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        left, right = 0, len(nums)-1
        if nums[left] <= nums[right]:
            return nums[left]

        while left < right:
            mid = (left + right) // 2
            if nums[mid] > nums[mid + 1]:
                return nums[mid + 1]
            if nums[mid - 1] > nums[mid]:
                return nums[mid]

            if nums[mid] > nums[0]:
                left = mid + 1
            else:
                right = mid - 1
        return -1

    def test(self):
        for nums, target in [
            ([1], 1),
        ]:
            ans = self.findMin(nums)
            assert ans == target, f"target {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
