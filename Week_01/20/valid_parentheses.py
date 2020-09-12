class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {
            "}": "{",
            "]": "[",
            ")": "(",
        }
        stack = ["?"]

        for c in s:
            if c not in pairs:
                stack.append(c)
            elif stack.pop() != pairs[c]:
                return False
        return len(stack) == 1

    def test(self):
        for s, target in [
            ("{", False),
            ("{[", False),
            ("{}", True),
            ("{()}", True),
            ("{()))", False),
        ]:
            assert target is self.isValid(s), f"failed. s: {s}"


if __name__ == "__main__":
    Solution().test()
