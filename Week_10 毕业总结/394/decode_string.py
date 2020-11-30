class Solution:
    def decodeString(self, s: str) -> str:
        stack = []
        i = 0

        def get_digits(p):
            n = ""
            while '0' <= s[p] <= '9':
                n += s[p]
                p += 1
            return n

        while i < len(s):
            if '0' <= s[i] <= '9':
                digits = get_digits(i)
                i += len(digits)
                stack.append(digits)
            elif s[i] != ']':
                stack.append(s[i])
                i += 1
            else:
                i += 1
                sub = ""
                while stack[-1] != "[":
                    sub = stack.pop() + sub
                stack.pop()
                repeat_time = int(stack.pop())
                stack.append(sub * repeat_time)

        return "".join(stack)
