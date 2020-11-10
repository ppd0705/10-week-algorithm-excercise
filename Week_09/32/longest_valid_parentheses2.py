class Solution:
    def longestValidParentheses(self, s: str) -> int:
        ans = 0
        stack = [-1]
        for i in range(len(s)):
            if s[i] == '(':
                stack.append(i)
            else:
                stack.pop()
                if len(stack) == 0:
                    stack.append(i)
                else:
                    ans = max(ans, i-stack[-1])
        return ans

    def test(self):
        s = "(())"
        ans = self.longestValidParentheses(s)
        target = 4
        assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
