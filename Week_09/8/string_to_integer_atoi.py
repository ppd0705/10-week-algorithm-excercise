class Solution:
    def myAtoi(self, str: str) -> int:
        res = 0
        sign = 1
        idx = 0

        while idx < len(str) and str[idx] == ' ':
            idx += 1

        if idx == len(str):
            return 0
        if str[idx] == '-':
            sign = -1
            idx += 1
        elif str[idx] == '+':
            idx += 1

        while idx < len(str) and '0' <= str[idx] <= '9':
            if sign == 1:
                if res > 2 ** 31 // 10 or res == 2 ** 31 // 10 and str[idx] >= '7':
                    return 2 ** 31 - 1
                res = res * 10 + ord(str[idx]) - ord('0')
            else:
                if res < -2 ** 31 // 10 or res == -(2 ** 31 // 10) and str[idx] >= '8':
                    return -2 ** 31
                res = res * 10 - ord(str[idx]) + ord('0')

            idx += 1
        return res
