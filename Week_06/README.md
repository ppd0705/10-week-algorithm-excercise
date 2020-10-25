学习笔记

|#|Title|Solutions|
|---|---|------|
|1143|[longest-common-subsequence](https://leetcode-cn.com/problems/longest-common-subsequence) | 递归([Go](1143/longest_common_subsequence.go),[Py](1143/longest_common_subsequence.py)),动态规划([Go](1143/longest_common_subsequence2.go),[Py](1143/longest_common_subsequence2.py)),动态规划2([Go](1143/longest_common_subsequence3.go),[Py](1143/longest_common_subsequence3.py))|
|213|[house-robber-ii](https://leetcode-cn.com/problems/house-robber-ii) | 动态规划([Go](213/house_robber_ii.go),[Py](213/house_robber_ii.go))|
|62|[unique-paths](https://leetcode-cn.com/problems/unique-paths) | 动态规划([Go](62/unique_paths.go),[Py](62/unique_paths.go))|
|63|[unique-paths-ii](https://leetcode-cn.com/problems/unique-paths-ii) | 动态规划([Go](63/unique_path_ii.go),[Py](63/unique_path_ii.go))|
|120|[triangle](https://leetcode-cn.com/problems/triangle) | 动态规划([Go](120/triangle.go),[Py](120/triangle.go))|
|53|[maximum-subarray](https://leetcode-cn.com/problems/maximum-subarray) | 动态规划([Go](../Week_06/53/maximum_subarray.go),[Py](../Week_06/53/maximum_subarray.py)),分治([Go](../Week_06/53/maximum_subarray2.go),[Py](../Week_06/53/maximum_subarray2.py))|
|152|[maximum-product-subarray](https://leetcode-cn.com/problems/maximum-product-subarray) | 动态规划([Go](../Week_06/152/maximum_product_subarray.go),[Py](../Week_06/152/maximum_product_subarray.py))|
|322|[coin-change](https://leetcode-cn.com/problems/coin-change) | 动态规划([Go](322/coin_change.go),[Py](322/coin_change.py))|
|198|[house-robber](https://leetcode-cn.com/problems/house-robber) | 动态规划([Go](198/house_robber.go),[Py](198/house_robber.py))|
|121|[best-time-to-buy-and-sell-stock](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) | 动态规划([Go](121/best_time_to_buy_and_sell_stock.go),[Py](121/best_time_to_buy_and_sell_stock.py))|

## 题解

### 1143. longest-common-subsequence

1. 递归 
2. 动态规划:
  - dp[i][j] = dp[i-1][j-1] + 1 或 max(dp[i][j-1], dp[i-1][j])
3. 动态规划空间优化: 
  - 用一个变量pre保存dp[i-1][j-1]可以实现空间负载度降维
  - dp[j] = pre+1 或者 max(dp[j], dp[j-1])
  
  

### 62. unique-paths

1. 动态规划:
  - dp[i][j] = 1 if (i==0 or j ==0) else dp[i-1][j] + dp[i][j-1]
  
  
### 120. triangle

1. 动态规划: 倒序递推:


### 53. maximum-subarray

1. 动态规划: dp[i] = nums[i] + max(0, dp[i-1])
    
2. 分治:


### 152. maximum-product-subarray

1. 动态规划: 用两个变量分别记录当前最小乘积和最大乘积

### 322. coin_change

1. 动态规划: 建立一个极值数组: dp = [amount+1] *(amount+1); dp[i] = min(dp[i], dp[i-c]+1)


### 198. house-robber

1. 动态规划:

### 213. house-robber-ii

1. 动态规划:
  - 两次动态规划dp(nums[:-1])、dp(nums[1:]),取两者之间的较大值
  
  

  
