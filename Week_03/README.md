学习笔记

|#|Title|Solutions|
|---|---|------|
|70|[climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs) | 递归([Go](70/climbing_stairs.go),[Py](70/climbing_stairs.py))|
|22|[generate-parentheses](https://leetcode-cn.com/problems/generate-parentheses) | 递归([Go](22/generate_parentheses.go),[Py](22/generate_parentheses.py))|
|226|[invert-binary-tree](https://leetcode-cn.com/problems/invert-binary-tree) | 递归([Go](226/invert_binary_tree.go),[Py](226/invert_binary_tree.go)),队列([Go](226/invert_binary_tree2.go))|
|98|[validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree) | 递归([Go](98/validate_binary_search_tree.go),[Py](98/validate_binary_search_tree.py)),中序遍历([Go](98/validate_binary_search_tree2.go),[Py](98/validate_binary_search_tree2.py))|
|104|[maximum-depth-of-binary-tree](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) | 递归([Go](104/maximum_depth_of_binary_tree.go)))|
|111|[minimum-depth-of-binary-tree](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 递归([Go](111/minimum_depth_of_binary_tree.go),[Py](111/minimum_depth_of_binary_tree.py)),层序遍历([Go](111/minimum_depth_of_binary_tree2.go),[Py](111/minimum_depth_of_binary_tree2.py))|
|236|[lowest-common-ancestor-of-a-binary-tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) | 递归([Go](236/lowest_common_ancestor_of_a-binary_tree.go),[Py](236/lowest_common_ancestor_of_a-binary_tree.py)),遍历记录父节点([Go](236/lowest_common_ancestor_of_a-binary_tree2.go),[Py](236/lowest_common_ancestor_of_a-binary_tree2.py))|


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

### 111. minimum-depth-of-binary-tree

1. 递归
2. 使用队列逐层遍历

### 236. lowest-common-ancestor-of-a-binary-tree

1. 递归
    - 如果root是p或q,则最近公共祖先就是root
    - 后续遍历
      - 如果left和right都有找到，那就是root
      - 如果left没找到，那返回right(right可能时nil)；right没找到亦然
2. 建2个哈希表，一个记录父节点，另一个记录是否访问过