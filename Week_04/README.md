学习笔记

|#|Title|Solutions|
|---|---|------|
|102|[binary-tree-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) | 广度优先([Go](102/binary_tree_level_order_traversal.go),[Py](102/binary_tree_level_order_traversal.py)),深度优先([Go](102/binary_tree_level_order_traversal2.go),[Py](102/binary_tree_level_order_traversal2.py))|
|433|[minimum-genetic-mutation](https://leetcode-cn.com/problems/minimum-genetic-mutation) | 广度优先([Go](433/minimum_genetic_mutation.go),[Py](433/minimum_genetic_mutation.py))|
|22|[generate-parentheses](https://leetcode-cn.com/problems/generate-parentheses) | 深度优先([Go](22/generate_parentheses.go),[Py](22/generate_parentheses.go))|
|122|[best-time-to-buy-and-sell-stock-ii](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii) | 贪心([Go](122/best_time_to_buy_and_sell_stock_ii.go),[Py](122/best_time_to_buy_and_sell_stock_ii.py))|
|860|[lemonade-change](https://leetcode-cn.com/problems/lemonade-change) | 贪心([Go](860/lemonade_change.go),[Py](860/lemonade_change.py))|
|515|[find-largest-value-in-each-tree-row](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row) | 广度优先([Go](515/find_largest_value_in_each_tree_row.go),[Py](515/find_largest_value_in_each_tree_row.py)),深度优先([Go](515/find_largest_value_in_each_tree_row2.go),[Py](515/find_largest_value_in_each_tree_row2.py))|
|127|[word-ladder](https://leetcode-cn.com/problems/word-ladder) | 广度优先([Go](127/word_ladder.go)),双向广度优先([Go](127/word_ladder2.go),[Py](127/word_ladder2.py))|
|126|[word-ladder-ii](https://leetcode-cn.com/problems/word-ladder-ii) | 广度优先+深度优先([Go](126/word_ladder_ii.go),[Py](126/word_ladder_ii.py))|
|200|[numbers-of-islands](https://leetcode-cn.com/problems/numbers-of-islands) | 深度优先([Go](200/number_of_islands.go),[Py](200/number_of_islands.py))|
|529|[minesweeper](https://leetcode-cn.com/problems/minesweeper) | 广度优先([Go](529/minesweeper.go),[Py](529/minesweeper.py)),深度优先([Go](529/minesweeper2.go),[Py](529/minesweeper2.py))|
|367|[valid-perfect-square](https://leetcode-cn.com/problems/valid-perfect-square) | 二分查找([Go](367/valid_perfect_square.go),[Py](367/valid_perfect_square.go))|
|322|[coin-change](https://leetcode-cn.com/problems/coin-change) | 递推([Go](322/coin_change.go),[Py](322/coin_change.py)),递归([Go](322/coin_change2.go),[Py](322/coin_change2.py))|
|455|[assign-cookies](https://leetcode-cn.com/problems/assign-cookies) | 贪心([Go](455/assign_cookies.go),[Py](455/assign_cookies.py))|
|874|[walking-robot-simulation](https://leetcode-cn.com/problems/walking-robot-simulation) | 迭代([Go](874/walking_robot_simulation.go),[Py](874/walking_robot_simulation.py))|
|55|[jump-game](https://leetcode-cn.com/problems/jump-game) | 贪心([Go](55/jump_game.go),[Py](55/jump_game.py))|
|45|[jump-game-ii](https://leetcode-cn.com/problems/jump-game-ii) | 贪心([Go](45/jump_game_ii.go),[Py](45/jump_game_ii.py))|
|69|[sqrtx](https://leetcode-cn.com/problems/sqrtx) | 二分查找([Go](69/sqrtx.go),[Py](69/sqrtx.py))|
|33|[search-in-rotated-sorted-array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) | 二分查找([Go](33/search_in_rotated_sorted_array.go),[Py](33/search_in_rotated_sorted_array.py))|
|74|[search-a-2d-matrix](https://leetcode-cn.com/problems/search-a-2d-matrix) | 二分查找([Go](74/search_a_2d_matrix.go),[Py](74/search_a_2d_matrix.py))|


## 题解

### 102. binary-tree-level-order-traversal

1. 队列: 广度优先，逐层入队出队
2. 递归: 深度优先


### 433. minimum-genetic-mutation

1. 广度优先，逐层入队出队

### 22. generate-parentheses

1. 递归：
  - dfs(left, right)
  - left < n 或 right < left 都是可能结果

### 860. lemonade-change

1. 贪心: 给20找零的时候优先找10+5而非5+5+5的组合

### 515. find-largest-value-in-each-tree-row

1. 广度优先: 层序遍历
2. 深度优先

### 127. word-ladder

1. 广度优先
2. 双向广度优先

### 126. word-ladder-ii

1. 广度优先+深度优先
    - 先将所有word按regex分组
    - 广度优先遍历，一次记录每个node最短路径的上一步nodes, queue空或者当前level+1> level_map[end_word]为终止条件
    - 深度优先反向遍历pre_nodes拼出完整的路径
    
### 200. numbers-of-islands

1. 深度优先

### 529. minesweeper

1. 广度优先: 
  - 层序遍历，依次遍历八个方向是否有'M'
  - 需要使用set去重
2. 深度优先


### 367. valid-perfect-square

1. 二分查找: for left <= right

### 322. coin-change

1. 递推: dp[i] = min(dp[i-c] for c in coins) + 1;return dp[amount]
2. 递归: 使用字典缓存子问题的结果

### 455. assign-cookies

1. 贪心: 先排序后双指针迭代

### 874. walking-robot-simulation

1. 迭代
 - 找到转向时dx/dy的变换规律
 - 一步一步往前走，遇到障碍回退一步即可 
 - 需要找出所有经过的点的距离最大值
 
### 55. jump-game

1. 贪心: 设置一个right_most的变量，初始值为0，小于right_most的点都能访问到，不断扩张right_most

### 45. jump-game-ii

1. 贪心: 
  - 设置边界值end初始值为0
  - 设置right_most，初值值为0
  - 当i==end时即到达边界，step++, end = right_most
  
### 69. sqrtx

1. 二分查找

### 33. search-in-rotated-sorted-array

1. 二分查找: 
  - 当mid < left是右半部分有序反之左半部分有序
  
### 74. search-a-2d-matrix

1. 二分查找: 
  - 先二分找到行
  - 再二分找列 
 
### 153. find-minimum-in-rotated-sorted-array

1. 二分查找1: 
    - nums[mid] > nums[right]则左半区递减，left=mid+1反之right=mid
2. 二分查找2:   
    - 先判断是否全局有序，如果是则返回nums[0]
    - 再判断
       - 当nums[mid] > nums[mid+1]返回mid+1
       - 当nums[mid] < nums[mid-1]返回mid
       - 都不满足时如果nums[mid] > nums[0]时最小值一点在右半区，否则在左半区
    