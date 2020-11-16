class Solution:
    def reverseOnlyLetters(self, S: str) -> str:
        left, right = 0, len(S) - 1
        chars = list(S)
        while left < right:
            if not 'A' <= chars[left] <= 'Z' and not 'a' <= chars[left] <= 'z':
                left += 1
            elif not 'A' <= chars[right] <= 'Z' and not 'a' <= chars[right] <= 'z':
                right -= 1
            else:
                chars[left], chars[right] = chars[right], chars[left]
                left += 1
                right -= 1
        return "".join(chars)
