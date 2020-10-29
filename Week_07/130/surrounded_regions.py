from typing import List


class Solution:
    def solve(self, board: List[List[str]]) -> None:
        if not board or not board[0]:
            return
        m, n = len(board), len(board[0])

        def dfs(x: int, y: int):
            if x < 0 or x >= m or y < 0 or y >= n or board[x][y] != "O":
                return
            board[x][y] = "B"
            dfs(x - 1, y)
            dfs(x + 1, y)
            dfs(x, y - 1)
            dfs(x, y + 1)

        for i in range(m):
            for j in range(n):
                if board[i][j] == "O" and i == 0 or i == m - 1 or j == 0 or j == n - 1:
                    dfs(i, j)
        for i in range(m):
            for j in range(n):
                if board[i][j] == "O":
                    board[i][j] = 'X'
                elif board[i][j] == "B":
                    board[i][j] = "O"
