class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        chars = list(s)
        i = 0
        n = len(s)
        while i < n:
            start = i
            end = min(i + k - 1, n - 1)
            while start < end:
                chars[start], chars[end] = chars[end], end[start]
                start += 1
                end -= 1
            i += 2 * k
        return str(chars)
