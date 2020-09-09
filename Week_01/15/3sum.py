from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        for i, a in enumerate(nums[:-2]):
            if a > 0:
                break
            if i > 0 and nums[i - 1] == a:
                continue
            l, r = i + 1, len(nums) - 1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if s == 0:
                    res.append([nums[i], nums[l], nums[r]])
                    l += 1
                    r -= 1
                    while l < r and nums[l] == nums[l - 1]:
                        l += 1
                    while l < r and nums[r] == nums[r + 1]:
                        r -= 1
                elif s < 0:
                    l += 1
                    while l < r and nums[l] == nums[l - 1]:
                        l += 1
                else:
                    r -= 1
                    while l < r and nums[r] == nums[r + 1]:
                        r -= 1
        return res

    def test(self):
        for nums, target in [
            ([1, 0, -1, -1], [[-1, 0, 1]]),
        ]:
            ans = self.threeSum(nums)
            assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
