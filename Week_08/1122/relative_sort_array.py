from typing import List


class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        counter = [0] * 1001

        for a in arr1:
            counter[a] += 1

        idx = 0
        for a in arr2:
            while counter[a] > 0:
                arr1[idx] = a
                idx += 1
                counter[a] -= 1
        for i, c in enumerate(counter):
            for j in range(c):
                arr1[idx + j] = i
            idx += c
        return arr1
