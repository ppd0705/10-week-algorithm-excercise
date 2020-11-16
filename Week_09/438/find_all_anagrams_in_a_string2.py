from typing import List


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        counter = [0] * 26
        for c in p:
            counter[ord(c) - ord('a')] += 1
        res = []

        left, right = 0, 0
        tmp_counter = [0] * 26
        while right < len(s):
            tmp_counter[ord(s[right]) - ord('a')] += 1
            while tmp_counter[ord(s[right]) - ord('a')] > counter[ord(s[right]) - ord('a')]:
                tmp_counter[ord(s[left]) - ord('a')] -= 1
                left += 1
            if right - left + 1 == len(p):
                res.append(left)
            right += 1
        return res
