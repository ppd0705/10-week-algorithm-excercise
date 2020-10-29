学习笔记

|#|Title|Solutions|
|---|---|------|
|208|[implement-trie-prefix-tree](https://leetcode-cn.com/problems/implement-trie-prefix-tree) | 前缀树([Go](../Week_07/208/implement_trie_prefix_tree.go),[Py](../Week_07/208/implement_trie_prefix_tree.py))|
|212|[word-search-ii](https://leetcode-cn.com/problems/word-search-ii) | 前缀树+回溯([Go](../Week_07/212/word_search_ii.go),[Py](../Week_07/212/word_search_ii.py))|
|547|[friend-circles](https://leetcode-cn.com/problems/friend-circles) | 深度优先([Go](../Week_07/547/friend_circles.go),[Py](../Week_07/547/friend_circles.py)),并查集([Go](../Week_07/547/friend_circles2.go),[Py](../Week_07/547/friend_circles2.py))|
|130|[surrounded-regions](https://leetcode-cn.com/problems/surrounded-regions) | 深度优先([Go](../Week_07/130/surrounded_regions.go),[Py](../Week_07/130/surrounded_regions.py)),并查集([Go](../Week_07/130/surrounded_regions2.go))|
|36|[valid-sudoku](https://leetcode-cn.com/problems/valid-sudoku) | 迭代([Go](../Week_07/36/valid_sudoku.go),[Py](../Week_07/36/valid_sudoku.py))|
|37|[sudoku-solver](https://leetcode-cn.com/problems/sudoku-solver) | 回溯([Go](../Week_07/37/sudoku_solver.go),[Py](../Week_07/37/sudoku_solver.py))|
|1091|[shortest-path-in-binary-matrix](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix) | 广度优先搜索([Go](../Week_07/1091/shortest_path_in_binary_matrix.go),[Py](../Week_07/1091/shortest_path_in_binary_matrix.py))|


## 题解

### 208. implement-trie-prefix-tree

1. 哈希表：使用哈希表记录前缀关系，用一个标志记录是否是叶子节点 

### 212. word-search-ii

1. 前缀树+回溯 
 
### 547. friend-circles

1. 深度优先搜索
2. 并查集
  
  
### 130. surrounded-regions

1. 深度优先搜索: 从边界上为"O"的点深度优先搜索
2. 并查集

### 36. valid-sudoku

1. 迭代：使用三个9*9的数组记录每行、列、块是否合法，Golang可以使用`[9]uint16`压缩空间

### 37. sudoku-solver

1. 回溯


### 1091. shortest-path-in-binary-matrix

1. 广度优先搜索