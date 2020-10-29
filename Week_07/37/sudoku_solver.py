from typing import List


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        row = [[False] * 9 for _ in range(9)]
        column = [[False] * 9 for _ in range(9)]
        box = [[False] * 9 for _ in range(9)]
        need_fill = [[False] * 9 for _ in range(9)]
        need_fill_count = 0

        def box_index(x: int, y: int) -> int:
            return x // 3 * 3 + y // 3

        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    need_fill[i][j] = True
                    need_fill_count += 1
                else:
                    v = int(board[i][j]) - 1
                    row[i][v] = True
                    column[j][v] = True
                    box[box_index(i, j)][v] = True

        def get_next(x: int, y: int) -> (int, int):
            if y == 8:
                return x + 1, 0
            return x, y + 1

        def dfs(x: int, y: int, cnt: int) -> bool:
            if cnt == 0:
                return True

            if not need_fill[x][y]:
                return dfs(*get_next(x, y), cnt)
            else:
                for v in range(9):
                    if not row[x][v] and not column[y][v] and not box[box_index(x, y)][v]:
                        row[x][v] = True
                        column[y][v] = True
                        box[box_index(x, y)][v] = True
                        need_fill[x][y] = False
                        board[x][y] = str(v + 1)
                        if dfs(*get_next(x, y), cnt - 1):
                            return True
                        row[x][v] = False
                        column[y][v] = False
                        box[box_index(x, y)][v] = False
                        need_fill[x][y] = True
                        board[x][y] = "."
                return False

        dfs(0, 0, need_fill_count)

    def test(self):
        board = [["5", "3", ".", ".", "7", ".", ".", ".", "."],
                 ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                 [".", "9", "8", ".", ".", ".", ".", "6", "."],
                 ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                 ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                 ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                 [".", "6", ".", ".", ".", ".", "2", "8", "."],
                 [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                 [".", ".", ".", ".", "8", ".", ".", "7", "9"]]

        target = [['5', '3', '4', '6', '7', '8', '9', '1', '2'],
                  ['6', '7', '2', '1', '9', '5', '3', '4', '8'],
                  ['1', '9', '8', '3', '4', '2', '5', '6', '7'],
                  ['8', '5', '9', '7', '6', '1', '4', '2', '3'],
                  ['4', '2', '6', '8', '5', '3', '7', '9', '1'],
                  ['7', '1', '3', '9', '2', '4', '8', '5', '6'],
                  ['9', '6', '1', '5', '3', '7', '2', '8', '4'],
                  ['2', '8', '7', '4', '1', '9', '6', '3', '5'],
                  ['3', '4', '5', '2', '8', '6', '1', '7', '9']]

        self.solveSudoku(board)

        assert board == target


if __name__ == "__main__":
    Solution().test()
