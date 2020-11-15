class Solution:
    def toLowerCase(self, str: str) -> str:
        return "".join(chr(ord(s)+32) if "A" <= s <= "Z" else s  for s in str)