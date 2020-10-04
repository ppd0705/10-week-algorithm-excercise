from typing import List


class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        x, y = click[0], click[1]
        row, column = len(board), len(board[0])
        if board[x][y] == 'M':
            board[x][y] = 'X'
            return board

        point_set = {(x, y)}

        while point_set:
            new_set = set()

            for (x, y) in point_set:
                pointers = []
                count = 0
                for dx, dy in [(-1, -1), (-1, 1), (-1, 0), (1, -1), (1, 1), (1, 0), (0, 1), (0, -1)]:
                    if 0 <= x + dx < row and 0 <= y + dy < column:
                        if board[x + dx][y + dy] == 'M':
                            count += 1
                        elif board[x + dx][y + dy] == 'E':
                            pointers.append((x + dx, y + dy))
                if count > 0:
                    board[x][y] = str(count)
                else:
                    board[x][y] = 'B'
                    new_set.update(pointers)
            point_set = new_set
        return board

