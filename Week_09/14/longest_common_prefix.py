from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        ans = strs[0]
        idx = 0
        flag = True

        while flag and idx < len(ans):
            for s in strs[1:]:
                if idx > len(s) or s[idx] != ans[idx]:
                    flag = False
                    break
            else:
                idx += 1
        return ans[:idx]
