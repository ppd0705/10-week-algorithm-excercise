class Solution:
    def removeOuterParentheses(self, S: str) -> str:
        counter = 0
        start = 0

        strs = []
        for i, c in enumerate(S):
            if c == "(":
                counter += 1
            else:
                counter -= 1

            if counter == 0:
                strs.append(S[start + 1:i])
                start = i + 1
        return "".join(strs)
