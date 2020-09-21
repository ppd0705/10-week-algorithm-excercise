from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        solution = [""] * 2 * n

        def dfs(left: int, right: int):
            if right == n:
                res.append("".join(solution))
                return
            if left < n:
                solution[left + right] = "("
                dfs(left + 1, right)
            if right < left:
                solution[left + right] = ")"
                dfs(left, right + 1)

        dfs(0, 0)
        return res
