from typing import List


class Solution:
    def slidingPuzzle(self, board: List[List[int]]) -> int:
        neighbors = {
            0: (1, 3),
            1: (0, 2, 4),
            2: (1, 5),
            3: (0, 4),
            4: (1, 3, 5),
            5: (2, 4),
        }

        curr = tuple(c for l in board for c in l)
        target = (1, 2, 3, 4, 5, 0)
        zero_index = curr.index(0)

        queue = [(curr, zero_index)]
        visited = {curr}

        res = -1
        while queue:
            res += 1
            l = len(queue)
            for i in range(l):
                curr, zero_index = queue[i]
                if curr == target:
                    return res
                curr_list = list(curr)
                for n in neighbors[zero_index]:
                    curr_list[zero_index], curr_list[n] = curr_list[n], curr_list[zero_index]
                    new = tuple(curr_list)
                    if new not in visited:
                        visited.add(new)
                        queue.append((new, n))
                    curr_list[zero_index], curr_list[n] = curr_list[n], curr_list[zero_index]
            queue = queue[l:]

        return -1
