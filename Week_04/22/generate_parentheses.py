from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        template = [""] * 2 * n
        res = []

        def dfs(left: int, right: int):
            if right == n:
                res.append("".join(template))
                return

            if left < n:
                template[left + right] = "("
                dfs(left + 1, right)
            if right < left:
                template[left + right] = ")"
                dfs(left, right + 1)

        if n > 0:
            dfs(0, 0)
        return res
