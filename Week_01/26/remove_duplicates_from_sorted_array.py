from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if not nums:
            return 0
        i = 0
        for j in range(1, len(nums)):
            if nums[j] != nums[i]:
                i += 1
                nums[i] = nums[j]
        return i + 1

    def test(self):
        for nums, target in [
            ([], 0),
            ([0], 1),
            ([0, 0], 1),
            ([0, 0, 1], 2),
            ([0, 0, 1, 2], 3),
            ([0, 0, 1, 2, 2, 2, 3, 4, 4], 5),
        ]:
            ans = self.removeDuplicates(nums)
            assert ans == target, f"target: {target}, ans: {ans}"

        print("all done")


if __name__ == "__main__":
    Solution().test()
