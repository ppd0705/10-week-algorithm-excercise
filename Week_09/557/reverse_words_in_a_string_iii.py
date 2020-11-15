class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join((c[::-1] for c in s.split()))

