from typing import List


class Solution:
    def getLeastNumbers(self, arr: List[int], k: int) -> List[int]:
        if k == 0:
            return []
        n = len(arr)
        if n == 0 or n == k:
            return arr

        for i in range(n // 2 - 1, -1, -1):
            self.heapify(arr, i, k - 1)

        for i in range(k, n):
            if arr[i] < arr[0]:
                arr[0] = arr[i]
                self.heapify(arr, 0, k - 1)
        return arr[:k]

    def heapify(self, arr: List[int], root: int, end: int):
        while True:
            child = root * 2 + 1
            if child > end:
                return

            if child + 1 <= end and arr[child + 1] > arr[child]:
                child += 1

            if arr[root] > arr[child]:
                return

            arr[root], arr[child] = arr[child], arr[root]
            root = child

    def test(self):

        for arr, k, target in [
            ([1, 4, 3, 4], 3, [1, 3, 4]),
            ([1, 2, 3, 4], 1, [1]),
            ([1, 2, 3, 4], 1, [1]),
            ([1, 4, 3, 2], 2, [1, 2]),
            ([1, 2, 3, 4], 3, [1, 2, 3]),
            ([1, 2, 3, 4], 4, [1, 2, 3, 4]),
        ]:
            ans = sorted(self.getLeastNumbers(arr, k))
            assert ans == target, f"ans: {ans}, target: {target}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
