学习笔记


|#|Title|Solutions|
|---|---|------|
|32|[longest-valid-parentheses](https://leetcode-cn.com/problems/longest-valid-parentheses) | 动态规划([Go](../Week_09/32/longest_valid_parentheses.go),[Py](../Week_09/32/longest_valid_parentheses.py)),栈([Go](../Week_09/32/longest_valid_parentheses2.go),[Py](../Week_09/32/longest_valid_parentheses2.py)),贪心([Go](../Week_09/32/longest_valid_parentheses3.go),[Py](../Week_09/32/longest_valid_parentheses3.py))|
|64|[minimum-path-sum](https://leetcode-cn.com/problems/minimum-path-sum) | 动态规划([Go](64/minimum_path_sum.go),[Py](64/minimum_path_sum.py))|
|72|[edit-distance](https://leetcode-cn.com/problems/edit-distance) | 动态规划([Go](72/edit_distance.go),[Py](72/edit_distance.py))|
|85|[maximal-rectangle](https://leetcode-cn.com/problems/maximal-rectangle) | 动态规划([Go](85/maximal_rectangle.go),[Py](85/maximal_rectangle.py))|
|115|[distinct-subsequences](https://leetcode-cn.com/problems/distinct-subsequences) | 动态规划([Go](115/distinct_subsequences.go),[Py](115/distinct_subsequences.py))|
|58|[length-of-last-word](https://leetcode-cn.com/problems/length-of-last-word) | 逆序迭代([Go](58/length_of_last_word.go),[Py](58/length_of_last_word.py))|
|709|[to-lower-case](https://leetcode-cn.com/problems/to-lower-case) | 迭代([Go](709/to_lower_case.go),[Py](709/to_lower_case.py))|
|8|[string-to-integer-atoi](https://leetcode-cn.com/problems/string-to-integer-atoi) | 迭代([Go](8/string_to_integer_atoi.go),[Py](8/string_to_integer_atoi.py))|
|14|[longest-common-prefix](https://leetcode-cn.com/problems/longest-common-prefix) | 迭代([Go](14/longest_common_prefix.go),[Py](14/longest_common_prefix.py))|
|344|[reverse-string](https://leetcode-cn.com/problems/reverse-string) | 双指针([Go](344/reverse_string.go),[Py](344/reverse_string.py))|
|541|[reverse-string-ii](https://leetcode-cn.com/problems/reverse-string-ii) | 迭代([Go](../Week_09/541/reverse_string_ii.go),[Py](../Week_09/541/reverse_string_ii.py))|
|557|[reverse-words-in-a-string-iii](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii) | 迭代([Go](../Week_09/557/reverse_words_in_a_string_iii.go),[Py](../Week_09/557/reverse_words_in_a_string_iii.py))|
|917|[reverse-only-letters](https://leetcode-cn.com/problems/reverse-only-letters) | 双指针([Go](../Week_09/917/reverse_only_letters.go),[Py](../Week_09/917/reverse_only_letters.py))|
|438|[find-all-anagrams-in-a-string](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string) | 暴力([Go](../Week_09/438/find_all_anagrams_in_a_string.go),滑动窗口([Go](../Week_09/438/find_all_anagrams_in_a_string2.go),[Py](../Week_09/438/find_all_anagrams_in_a_string2.py))|
|125|[valid-palindrome](https://leetcode-cn.com/problems/valid-palindrome) | 双指针([Go](../Week_09/125/valid_palindrome.go),[Py](../Week_09/125/valid_palindrome.py))|
|680|[valid-palindrome-ii](https://leetcode-cn.com/problems/valid-palindrome-ii) | 双指针夹逼([Go](../Week_09/680/valid_palindrome_ii.go),[Py](../Week_09/680/valid_palindrome_ii.py))|
|5|[longest-palindromic-substring](https://leetcode-cn.com/problems/longest-palindromic-substring) | 动态规划([Go](../Week_09/5/longest_palindromic_substring.go),[Py](../Week_09/5/longest_palindromic_substring.py)),中心扩展([Go](../Week_09/5/longest_palindromic_substring2.go),[Py](../Week_09/5/longest_palindromic_substring2.py))|



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
    
### 72. edit-distance

1. 动态规划：
    - 当word1[i-1] == word2[j-1]时 dp[i][j] = dp[i-1][j-1],否则 min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1
    
### 85. maximal-rectangle

1. 动态规划：
    - 使用三个数组left[n], right[n],height[n]分别记录当前的左右边界和高
    
### 115. distinct-subsequences

1. 动态规划：
    - 当s[i-1] != t[j-1]时： dp[i][j] = dp[i-1][j] (不取s[j]) 否则 dp[i][j] = dp[i-1][j] + dp[i-1][j-1]