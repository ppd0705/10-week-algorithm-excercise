class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        res = 0
        for i in range(len(s) - 1, -1, -1):
            if s[i] == ' ':
                if res == 0:
                    continue
                else:
                    break
            else:
                res += 1
        return res
