class Solution:
    def firstUniqChar(self, s: str) -> int:
        arr = [0] * 26

        for i, c in enumerate(s):
            arr[ord(c) - ord('a')] = i

        for i, c in enumerate(s):
            if arr[ord(c) - ord('a')] == i:
                return i
            arr[ord(c) - ord('a')] = -1
        return -1
