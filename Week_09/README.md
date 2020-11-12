学习笔记


|#|Title|Solutions|
|---|---|------|
|32|[longest-valid-parentheses](https://leetcode-cn.com/problems/longest-valid-parentheses) | 动态规划([Go](../Week_09/32/longest_valid_parentheses.go),[Py](../Week_09/32/longest_valid_parentheses.py)),栈([Go](../Week_09/32/longest_valid_parentheses2.go),[Py](../Week_09/32/longest_valid_parentheses2.py)),贪心([Go](../Week_09/32/longest_valid_parentheses3.go),[Py](../Week_09/32/longest_valid_parentheses3.py))|
|64|[minimum-path-sum](https://leetcode-cn.com/problems/minimum-path-sum) | 动态规划([Go](64/minimum_path_sum.go),[Py](64/minimum_path_sum.py))|
|72|[edit-distance](https://leetcode-cn.com/problems/edit-distance) | 动态规划([Go](72/edit_distance.go),[Py](72/edit_distance.py))|
|85|[maximal-rectangle](https://leetcode-cn.com/problems/maximal-rectangle) | 动态规划([Go](85/maximal_rectangle.go),[Py](85/maximal_rectangle.py))|
|115|[distinct-subsequences](https://leetcode-cn.com/problems/distinct-subsequences) | 动态规划([Go](115/distinct_subsequences.go),[Py](115/distinct_subsequences.py))|



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
    
    
## 题解

### 72. edit-distance

1. 动态规划：
    - 当word1[i-1] == word2[j-1]时 dp[i][j] = dp[i-1][j-1],否则 min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1
    
### 85. maximal-rectangle

1. 动态规划：
    - 使用三个数组left[n], right[n],height[n]分别记录当前的左右边界和高
    
### 85. distinct-subsequences

1. 动态规划：
    - 当s[i-1] != t[j-1]时： dp[i][j] = dp[i-1][j] (不取s[j]) 否则 dp[i][j] = dp[i-1][j] + dp[i-1][j-1]