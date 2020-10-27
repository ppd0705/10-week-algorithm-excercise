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
|91|[decode-ways](https://leetcode-cn.com/problems/decode-ways) | 动态规划([Go](91/decode_ways.go),[Py](91/decode_ways.py))|
|123|[best-time-to-buy-and-sell-stock-iii](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii) | 动态规划([Go](123/best_time_to_buy_and_sell_stock_iii.go),[Py](123/best_time_to_buy_and_sell_stock_iii.py)),动态规划2([Go](123/best_time_to_buy_and_sell_stock_iii_2.go),[Py](123/best_time_to_buy_and_sell_stock_iii_2.py))|
|309|[best-time-to-buy-and-sell-stock-with-cooldown](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown) | 动态规划([Go](309/best_time_to_buy_and_sell_stock_with_cooldown.go),[Py](309/best_time_to_buy_and_sell_stock_with_cooldown.py))|
|188|[best-time-to-buy-and-sell-stock-iv](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv) | 动态规划([Go](188/best_time_to_buy_and_sell_stock_iv_2.go),[Py](188/best_time_to_buy_and_sell_stock_iv_2.py))|
|714|[best-time-to-buy-and-sell-stock-with-transaction-fee](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee) | 动态规划([Go](714/best_time_to_buy_and_sell_stock_with_transaction_fee.go),[Py](714/best_time_to_buy_and_sell_stock_with_transaction_fee.py))|
|279|[perfect-squares](https://leetcode-cn.com/problems/perfect-squares) | 动态规划([Go](279/perfect_squares.go),[Py](279/perfect_squares.py))|
|518|[coin-change-2](https://leetcode-cn.com/problems/coin-change-2) | 动态规划([Go](518/coin_change_2.go),[Py](518/coin_change_2.py))|

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
  
  
### 91. decode-ways
1. 动态规划: 首先需要将这道题类比爬楼梯问题，可能只能走两步，可能只能走一步，可能一步两步都可能走
    - 特殊处理s[i] == '0'
       - 如果s[i-1] in ('1', '2')， 即走两步: dp[i+1] = dp[i-1]
       - 否则不合法
    - 当s[i-1] == '1' 或者 s[i-1] == '2' && '1' <= s[i] <= '6'， 即走一步或两步: dp[i+1] = dp[i] + dp[i-1]
    - 否则走一步 dp[i+1] = dp[i]


### 123. best-time-to-buy-and-sell-stock-iii
1. 动态规划: 
  - 定义一个三维数组：第i天第j次买(0)或者卖(1)
  - 递推公式
    - 第0天 dp[i][j][1] = -prices[i]
    - 第1~n天
      - dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
      - dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
      
2. 动态规划二：
  因为最多只有两次交易，所以只需要定义四个变量 dp11(第一次买入)、dp10(第一次卖出)、dp21(第二次买入)、dp20(第二次卖出)
  - 第0天：
     - dp11 = -prices[0]
     - dp10 = 0
     - dp21 = -prices[0] （当做第0天买入两次卖出一次）
     - dp20 = 0 （当做第0天买入两次卖出两次）
  - 第1~n天:
     - dp11 = max(dp11, -prices[i])
     - dp10 = max(dp10, dp11+prices[i])
     - dp21 = max(dp21, dp10-prices[i])
     - dp20 = max(dp20, dp21+prices[i])
     
     
### 309. best-time-to-buy-and-sell-stock-with-cooldown
1. 动态规划:
  - 定义一个二维数：第i天是否买(1)或卖(0)
  - 递推
    - 第0，1天
      - dp[0][1] = -prices[0]
      - dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
      - dp[1][1] = max(dp[0][1], dp[0][0]-prices[1])
    - 第2~n天
      - dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
      - dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

### 309. best-time-to-buy-and-sell-stock-iv
1. 动态规划: #123的通用情况

### 714. best-time-to-buy-and-sell-stock-with-transaction-fee
1. 动态规划: 每一天只有持有和不持有两个状态，且只和前一天有关,故可以用两个变量hold和not_hold即可
  - not_hold = max(not_hold, hold+prices[i])
  - hold = max(hold, not_hold-prices[i]-fee)


### 279. perfect-squares
1. 动态规划: 类似于换硬币问题

### 518. coin-change-2
1. 动态规划