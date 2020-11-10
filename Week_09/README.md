学习笔记


|#|Title|Solutions|
|---|---|------|
|32|[longest-valid-parentheses](https://leetcode-cn.com/problems/longest-valid-parentheses) | 动态规划([Go](../Week_09/32/longest_valid_parentheses.go),[Py](../Week_09/32/longest_valid_parentheses.py)),栈([Go](../Week_09/32/longest_valid_parentheses2.go),[Py](../Week_09/32/longest_valid_parentheses2.py)),贪心([Go](../Week_09/32/longest_valid_parentheses3.go),[Py](../Week_09/32/longest_valid_parentheses3.py))|



## 题解

### 32. longest-valid-parentheses

1. 动态规划：
    - 记录s[i] == ')'的dp[i]
      - 当s[i-1] == '(': dp[i] = dp[i-2] + 2
      - 当s[i-1] == '('且dp[i-1] > 0: dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
2. 栈: 
    - 用栈记录下标，初始加一个`-1`的值[-1]
    - 当s[i] == '('时，i入栈， 否则弹出栈顶元素，之后当前栈顶即为左边界:len = i - stack[len(stack)-1]
3: 贪心:
    - 用left和right分别记录左右括号的数量
    - 先从左向右遍历，遇到'('则left++,否则right++,当left==right为可能的解，right>left时将left/right重置为0
    - 从右向左遍历亦然
    