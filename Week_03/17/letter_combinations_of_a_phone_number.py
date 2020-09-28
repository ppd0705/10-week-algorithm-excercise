from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        number_letter_map = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        res = []
        base = [""] * len(digits)

        def dfs(idx: int):
            if idx == len(digits):
                res.append("".join(base))
                return
            for c in number_letter_map[digits[idx]]:
                base[idx] = c
                dfs(idx + 1)

        if digits:
            dfs(0)
        return res
