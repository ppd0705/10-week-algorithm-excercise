from typing import List


class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        x, y = click[0], click[1]
        row, column = len(board), len(board[0])
        if board[x][y] == 'M':
            board[x][y] = 'X'
            return board

        def dfs(r, c):
            count = 0
            for dx, dy in [(-1, -1), (-1, 1), (-1, 0), (1, -1), (1, 1), (1, 0), (0, 1), (0, -1)]:
                if 0 <= r + dx < row and 0 <= c + dy < column and board[r + dx][c + dy] == 'M':
                    count += 1
            if count > 0:
                board[r][c] = str(count)
            else:
                board[r][c] = 'B'
                for dx, dy in [(-1, -1), (-1, 1), (-1, 0), (1, -1), (1, 1), (1, 0), (0, 1), (0, -1)]:
                    if 0 <= r + dx < row and 0 <= c + dy < column and board[r + dx][c + dy] == 'E':
                        dfs(r + dx, c + dy)

        dfs(x, y)
        return board
