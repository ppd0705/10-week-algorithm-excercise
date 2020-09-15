class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        counter = [0] * 26

        for c in s:
            counter[ord(c) - ord('a')] += 1
        for c in t:
            counter[ord(c) - ord('a')] -= 1

        for i in counter:
            if i != 0:
                return False
        return True
