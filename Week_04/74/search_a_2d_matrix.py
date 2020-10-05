from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m = len(matrix)
        if m == 0:
            return False
        n = len(matrix[0])
        if n == 0:
            return False

        left, right = 0, m

        while left <= right:
            mid = (left + right) // 2
            if target < matrix[mid][0]:
                right = mid - 1
            elif target > matrix[mid][n - 1]:
                left = mid + 1
            else:
                l, r = 0, n - 1
                while l <= r:
                    c = (l + r) // 2
                    if matrix[mid][c] == target:
                        return True
                    elif target < matrix[mid][c]:
                        r = c - 1
                    else:
                        l = c + 1
                return False
        return False
