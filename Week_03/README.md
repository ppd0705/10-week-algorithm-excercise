学习笔记

|#|Title|Solutions|
|---|---|------|
|70|[climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs) | 递归([Go](70/climbing_stairs.go),[Py](70/climbing_stairs.py))|
|22|[generate-parentheses](https://leetcode-cn.com/problems/generate-parentheses) | 递归([Go](22/generate_parentheses.go),[Py](22/generate_parentheses.py))|
|226|[invert-binary-tree](https://leetcode-cn.com/problems/invert-binary-tree) | 递归([Go](226/invert_binary_tree.go),[Py](226/invert_binary_tree.go)),队列([Go](226/invert_binary_tree2.go))|
|98|[validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree) | 递归([Go](98/validate_binary_search_tree.go),[Py](98/validate_binary_search_tree.py)),中序遍历([Go](98/validate_binary_search_tree2.go),[Py](98/validate_binary_search_tree2.py))|


## 题解

### 70. climbing-stairs

除了递推外也可以用递归，不过涉及到重复计算，需要用哈希表做缓存

### 22. generate-parentheses

递归

### 226. invert-binary-tree

1. 递归
2. 队列

### 98. validate-binary-search-tree

1. 递归
2. 中序遍历