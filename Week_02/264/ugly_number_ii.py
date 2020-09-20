import heapq


class Solution:
    def nthUglyNumber(self, n: int) -> int:
        nums = []
        heapq.heappush(nums, 1)
        curr = 1
        seen = {1}
        for i in range(n):
            curr = heapq.heappop(nums)
            for f in (2, 3, 5):
                n = f * curr
                if n not in seen:
                    seen.add(n)
                    heapq.heappush(nums, n)

        return curr

    def test(self):
        for n, target in [
            (1, 1),
            (10, 12),
        ]:
            ans = self.nthUglyNumber(n)
            assert ans == target, f"n: {n}, target: {target}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
