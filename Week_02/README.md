学习笔记

|#|Title|Solutions|
|---|---|------|
|242|[valid-anagram](https://leetcode-cn.com/problems/anagram) | 数组计数([Go](242/valid_anagram.go),[Py](242/valid_anagram.py))|
|49|[group-anagrams](https://leetcode-cn.com/problems/group-anagrams) | 计数分类([Go](49/group_anagrams.go),[Py](49/group_anagrams.py))|
|94|[binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 递归([Go](94/binary_tree_inorder_traversal.go),[Py](94/binary_tree_inorder_traversal.py)),栈([Go](94/binary_tree_inorder_traversal2.go),[Py](94/binary_tree_inorder_traversal2.py)), Morris遍历([Go](94/binary_tree_inorder_traversal3.go),[Py](94/binary_tree_inorder_traversal3.py))|
|144|[binary-tree-preorder-traversal](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 递归([Go](144/binary_tree_preorder_traversal.go),[Py](144/binary_tree_preorder_traversal.py)),栈([Go](144/binary_tree_preorder_traversal2.go),[Py](144/binary_tree_preorder_traversal2.py)), Morris遍历([Go](144/binary_tree_preorder_traversal3.go),[Py](144/binary_tree_preorder_traversal3.py))|
|590|[n-ary-tree-postorder-traversal](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal) | 递归([Go](590/n_ary_tree_postorder_traversal.go),[Py](590/n_ary_tree_postorder_traversal.go)),栈([Go](590/n_ary_tree_postorder_traversal2.go),[Py](590/n_ary_tree_postorder_traversal2.go))|
|589|[n-ary-tree-preorder-traversal](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal) | 递归([Go](589/n_ary_tree_preorder_traversal.go),[Py](589/n_ary_tree_preorder_traversal.go)),栈([Go](589/n_ary_tree_preorder_traversal2.go),[Py](589/n_ary_tree_preorder_traversal2.go))|
|429|[n-ary-tree-level-order-traversal](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal) | 递归([Go](429/n_ary_tree_level_order_traversal.go),[Py](429/n_ary_tree_level_order_traversal.py)),栈([Go](429/n_ary_tree_level_order_traversal2.go),[Py](429/n_ary_tree_level_order_traversal2.py))|
|LCOF40|[top-k](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/) | 大顶堆([Go](LCOF40/least_numbers.go),[Py](LCOF40/least_numbers.py)),快速搜索([Go](LCOF40/least_numbers2.go),[Py](LCOF40/least_numbers2.py))|
|264|[ugly-number-ii/](https://leetcode-cn.com/problems/ugly-number-ii) | 小顶堆([Go](264/ugly_number_ii.go),[Py](264/ugly_number_ii.py)), 三指针([Go](264/ugly_number_ii2.go),[Py](264/ugly_number_ii2.py))|



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

### 590. n-ary-tree-postorder-traversal

1. 递归
2. 栈：先前序后翻转

### 589. n-ary-tree-preorder-traversal

1. 递归
2. 栈

### 429. n-ary-tree-level-order-traversal

1. 递归
2. 栈


### LCOF40. top-k

1. 大顶堆：取前k个数构建大顶堆，后续元素若小于heap[0],则替换之并shift down
2. 快速搜索：利用快排思想，当partition==k时即满足条件

### 264. ugly-number-ii

1. 小顶堆：依次弹出堆顶元素，如果之前没有入堆，则加入堆，第n次弹出的元素即为结果
2. 三指针：用三个指针依次记录2，3，5因子当前指向的丑数，