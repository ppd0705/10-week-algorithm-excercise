学习笔记

|#|Title|Solutions|
|---|---|------|
|242|[valid-anagram](https://leetcode-cn.com/problems/anagram) | 数组计数([Go](242/valid_anagram.go),[Py](242/valid_anagram.py))|
|49|[group-anagrams](https://leetcode-cn.com/problems/group-anagrams) | 计数分类([Go](49/group_anagrams.go),[Py](49/group_anagrams.py))|
|94|[binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 递归([Go](94/binary_tree_inorder_traversal.go),[Py](94/binary_tree_inorder_traversal.py)),栈([Go](94/binary_tree_inorder_traversal2.go),[Py](94/binary_tree_inorder_traversal2.py)), Morris遍历([Go](94/binary_tree_inorder_traversal3.go),[Py](94/binary_tree_inorder_traversal3.py))|
|144|[binary-tree-preorder-traversal](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 递归([Go](144/binary_tree_preorder_traversal.go),[Py](144/binary_tree_preorder_traversal.py)),栈([Go](144/binary_tree_preorder_traversal2.go),[Py](144/binary_tree_preorder_traversal2.py)), Morris遍历([Go](144/binary_tree_preorder_traversal3.go),[Py](144/binary_tree_preorder_traversal3.py))|


## 题解

### 242. valid-anagram

使用长度为26的数组计数，s++, t--


### 49. group-anagrams

使用长度为26的数组计数, 相同的即在同一组


### 94. binary-tree-inorder-traversal

1. 递归
2. 栈：将所有left和left.left先进栈，pop出最新的一个作为当前节点，再切换到当前节点的right
3. 如果无left,只直接切换到right,否则让root当left的最右子节点的右节点

### 144. binary-tree-prorder-traversal

1. 递归
2. 栈：right进栈
3. 如果无left,只直接切换到right,否则让right当left的最右子节点的右节点