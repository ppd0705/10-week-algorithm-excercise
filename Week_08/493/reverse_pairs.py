import bisect


class Solution:
    def reversePairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = []
        count = 0
        for i in range(len(nums) - 1, -1, -1):
            count += bisect.bisect_left(s, nums[i])
            bisect.insort_left(s, nums[i] * 2)
        return count


if __name__ == "__main__":
    Solution().reversePairs([2, 4, 3, 5, 1])
