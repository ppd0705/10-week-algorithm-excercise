学习笔记

|#|Title|Solutions|
|---|---|------|
|1143|[longest-common-subsequence](https://leetcode-cn.com/problems/longest-common-subsequence) | 递归([Go](1143/longest_common_subsequence.go),[Py](1143/longest_common_subsequence.py)),动态规划([Go](1143/longest_common_subsequence2.go),[Py](1143/longest_common_subsequence2.py)),动态规划2([Go](1143/longest_common_subsequence3.go),[Py](1143/longest_common_subsequence3.py))|
|213|[house-robber-ii](https://leetcode-cn.com/problems/house-robber-ii) | 动态规划([Go](213/house_robber_ii.go),[Py](213/house_robber_ii.go))|


## 题解

### 1143. longest-common-subsequence

1. 递归 
2. 动态规划:
  - dp[i][j] = dp[i-1][j-1] + 1 或 max(dp[i][j-1], dp[i-1][j])
3. 动态规划空间优化: 
  - 用一个变量pre保存dp[i-1][j-1]可以实现空间负载度降维
  - dp[j] = pre+1 或者 max(dp[j], dp[j-1])
  
  
### 213. house-robber-ii

1. 动态规划:
  - 两次动态规划dp(nums[:-1])、dp(nums[1:]),取两者之间的较大值