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
|200|[numbers-of-islands](https://leetcode-cn.com/problems/numbers-of-islands) | 深度优先([Go](200/number_of_islands.go),[Py](200/number_of_islands.py))|
|367|[valid-perfect-square](https://leetcode-cn.com/problems/valid-perfect-square) | 二分查找([Go](367/valid_perfect_square.go),[Py](367/valid_perfect_square.go))|


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

### 200. numbers-of-islands

1. 深度优先

### 367. valid-perfect-square

1. 二分查找: for left <= right