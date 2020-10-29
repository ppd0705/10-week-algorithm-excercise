from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row = [[False] * 9 for _ in range(9)]
        column = [[False] * 9 for _ in range(9)]
        box = [[False] * 9 for _ in range(9)]

        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    continue
                v = int(board[i][j]) - 1
                k = i // 3 * 3 + j // 3
                if row[i][v] or column[j][v] or box[k][v]:
                    return False
                row[i][v] = True
                column[j][v] = True
                box[k][v] = True
        return True
