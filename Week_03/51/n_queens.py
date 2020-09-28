from typing import List
from collections import defaultdict


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        row_pos = [0] * n
        column_used = defaultdict(bool)
        left_diagonal_used = defaultdict(bool)  # row + column
        right_diagonla_used = defaultdict(bool)  # row - column

        res = []

        def print_queue():
            template = [""] * n
            queue = []
            for i in range(n):
                for j in range(n):
                    if row_pos[i] == j:
                        template[j] = "Q"
                    else:
                        template[j] = "."
                queue.append("".join(template))
            return queue

        def dfs(row: int):
            if row == n:
                res.append(print_queue())
                return

            for col in range(n):
                if column_used[col] or left_diagonal_used[row + col] or right_diagonla_used[row - col]:
                    continue
                row_pos[row] = col
                column_used[col] = True
                left_diagonal_used[row + col] = True
                right_diagonla_used[row - col] = True
                dfs(row + 1)
                column_used[col] = False
                left_diagonal_used[row + col] = False
                right_diagonla_used[row - col] = False

        dfs(0)
        return res
