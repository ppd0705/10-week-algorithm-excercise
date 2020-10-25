from typing import List


class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        ans = max_product = min_product = nums[0]
        for num in nums[1:]:
            m, n = max_product, min_product
            max_product = max(m * num, n * num, num)
            min_product = min(m * num, n * num, num)
            ans = max(ans, max_product)
        return ans
