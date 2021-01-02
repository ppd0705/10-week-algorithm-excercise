
def heapify(nums, root, end):
    while True:
        child = root*2+1
        if child > end:
            return

        if child < end and nums[child+1] > nums[child]:
            child += 1

        if nums[child] <= nums[root]:
            return

        nums[root], nums[child] = nums[child], nums[root]
        root = child

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        n = len(nums)
        for i in range(n//2-1, -1, -1):
            heapify(nums, i, n-1)

        # pop k-1 times
        for i in range(k-1):
            nums[0], nums[n-i-1]= nums[n-i-1], nums[0]
            heapify(nums, 0, n-i-1-1)
        return nums[0]





