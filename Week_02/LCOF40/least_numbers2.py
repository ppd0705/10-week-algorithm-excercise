from typing import List


class Solution:
    def getLeastNumbers(self, arr: List[int], k: int) -> List[int]:
        if k == 0:
            return []
        n = len(arr)
        if n == 0 or n == k:
            return arr
        self.quick_search(arr, 0, n - 1, k)
        return arr[:k]

    def quick_search(self, arr: List[int], left: int, right: int, k: int):
        n = self.partition(arr, left, right)
        if n > k:
            return self.quick_search(arr, left, n - 1, k)
        elif n < k:
            return self.quick_search(arr, n+1, right, k)

    def partition(self, arr: List[int], left: int, right: int) -> int:
        pivot = arr[right]
        i = left - 1

        for j in range(left, right):
            if arr[j] < pivot:
                i += 1
                arr[i], arr[j] = arr[j], arr[i]
        i += 1
        arr[i], arr[right] = arr[right], arr[i]
        return i

    def test(self):

        for arr, k, target in [
            ([1, 4, 3, 4], 3, [1, 3, 4]),
            ([1, 2, 3, 4], 1, [1]),
            ([1, 2, 3, 4], 1, [1]),
            ([1, 4, 3, 2], 2, [1, 2]),
            ([1, 2, 3, 4], 3, [1, 2, 3]),
            ([1, 2, 3, 4], 4, [1, 2, 3, 4]),
        ]:
            print("nums:", arr)
            ans = sorted(self.getLeastNumbers(arr, k))
            assert ans == target, f"ans: {ans}, target: {target}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
