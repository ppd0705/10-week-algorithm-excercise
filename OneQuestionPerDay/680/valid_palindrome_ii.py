class Solution:
    def validPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1

        def helper(s, i, j):
            while i < j:
                if s[i] != s[j]:
                    return False
                i += 1
                j -= 1
            return True

        while left < right:
            if s[left] == s[right]:
                left += 1
                right -= 1
            else:
                return helper(s, left + 1, right) or helper(s, left, right - 1)
        return True
