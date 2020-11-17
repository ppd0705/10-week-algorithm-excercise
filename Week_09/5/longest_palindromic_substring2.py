class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) < 2:
            return s
        start, end = 0, 0

        for i in range(len(s)):
            l1 = self.center_spread(s, i, i)
            l2 = self.center_spread(s, i, i + 1)
            l = max(l1, l2)
            if l - 1 > end - start:
                start = i - (l - 1) // 2
                end = i + l // 2
        return s[start:end + 1]

    def center_spread(self, s: str, i: int, j: int) -> int:
        while i >= 0 and j < len(s) and s[i] == s[j]:
            i -= 1
            j += 1
        return j - i - 1
