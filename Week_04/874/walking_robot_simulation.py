from typing import List


class Solution:
    def robotSim(self, commands: List[int], obstacles: List[List[int]]) -> int:
        res = 0
        x, y, dx, dy = 0, 0, 0, 1
        obstacle_set = {
            (o[0], o[1]) for o in obstacles
        }
        for c in commands:
            if c == -2:
                dx, dy = -dy, dx
            elif c == -1:
                dx, dy = dy, -dx
            else:
                next_x, next_y = x, y
                for i in range(c):
                    next_x += dx
                    next_y += dy
                    if (next_x, next_y) in obstacle_set:
                        next_x -= dx
                        next_y -= dy
                        break
                x, y = next_x, next_y
                distance = x * x + y * y
                if distance > res:
                    res = distance
        return res
