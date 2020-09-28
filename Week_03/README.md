学习笔记

|#|Title|Solutions|
|---|---|------|
|70|[climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs) | 递归([Go](70/climbing_stairs.go),[Py](70/climbing_stairs.py))|
|22|[generate-parentheses](https://leetcode-cn.com/problems/generate-parentheses) | 递归([Go](22/generate_parentheses.go),[Py](22/generate_parentheses.py))|
|226|[invert-binary-tree](https://leetcode-cn.com/problems/invert-binary-tree) | 递归([Go](226/invert_binary_tree.go),[Py](226/invert_binary_tree.go)),队列([Go](226/invert_binary_tree2.go))|
|98|[validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree) | 递归([Go](98/validate_binary_search_tree.go),[Py](98/validate_binary_search_tree.py)),中序遍历([Go](98/validate_binary_search_tree2.go),[Py](98/validate_binary_search_tree2.py))|
|104|[maximum-depth-of-binary-tree](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) | 递归([Go](104/maximum_depth_of_binary_tree.go))|
|111|[minimum-depth-of-binary-tree](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 递归([Go](111/minimum_depth_of_binary_tree.go),[Py](111/minimum_depth_of_binary_tree.py)),层序遍历([Go](111/minimum_depth_of_binary_tree2.go),[Py](111/minimum_depth_of_binary_tree2.py))|
|236|[lowest-common-ancestor-of-a-binary-tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) | 递归([Go](236/lowest_common_ancestor_of_a-binary_tree.go),[Py](236/lowest_common_ancestor_of_a-binary_tree.py)),遍历记录父节点([Go](236/lowest_common_ancestor_of_a-binary_tree2.go),[Py](236/lowest_common_ancestor_of_a-binary_tree2.py))|
|297|[serialize-and-deserialize-binary-tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/) | 层序遍历([Go](297/serialize_and_deserialize_binary_tree.go),[Py](297/serialize_and_deserialize_binary_tree.py)),递归前序遍历([Go](297/serialize_and_deserialize_binary_tree2.go),[Py](297/serialize_and_deserialize_binary_tree2.py))|
|105|[construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal) | 递归([Go](105/construct_binary_tree_from_preorder_and_inorder_traversal.go),[Py](105/construct_binary_tree_from_preorder_and_inorder_traversal.py))|
|77|[combinations](https://leetcode-cn.com/problems/combinations) | 递归([Go](77/combinations.go),[Py](77/combinations.go))|
|46|[permutations](https://leetcode-cn.com/problems/permutations) | 递归([Go](46/permutations.go),[Py](46/permutations.py))|
|47|[permutations-ii](https://leetcode-cn.com/problems/permutations-ii) | 递归([Go](47/permutations.go),[Py](47/permutations.py))|
|50|[powx_n](https://leetcode-cn.com/problems/powx-n) | 递归([Go](50/powx_n.go),[Py](50/powx_n.py))|
|78|[subsets](https://leetcode-cn.com/problems/subsets) | 递归([Go](78/subsets.go),[Py](78/subsets.py))|
|169|[majority-element](https://leetcode-cn.com/problems/majority-element) | 哈希表([Go](169/majority_element.go),[Py](169/majority_element.py)),计数投票([Go](169/majority_element2.go),[Py](169/majority_element2.py))|
|17|[letter-combinations-of-a-phone-number](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number) | 递归([Go](17/letter_combinations_of_a_phone_number.go),[Py](17/letter_combinations_of_a_phone_number.go))|


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


### 297. serialize-and-deserialize-binary-tree/

1. 使用队列层序遍历
2. 递归前序遍历

### 297. construct-binary-tree-from-preorder-and-inorder-traversal

1. 递归
  - 前序遍历的第一个元素为root
  - 在中序遍历中找到root的index，从而得出左右子树的长度
  - 递归左右子树
  
### 77. combinations

1. 递归
  - 提前建立一个长度为k的空数组
  - 递归每次填一个数
  - 第k个即为结果
  
### 46. permutations

1. 递归+回溯
  - 提前建立一个长度为k的空数组,并建立一个同样长度为k的空数组用于标识是否已被使用
  - 递归每次填一个数
  - 第k个即为结果
  
### 47. permutations-ii

1. 递归+回溯
  - 提前建立一个长度为k的空数组,并建立一个字典记录当前每个元素剩余的个数
  - 递归每次填一个数
  - 第k个即为结果
  
  
### 50. powx-n

1. 递归
 - n为0时返回1
 - n<0时：n=-n,x = 1/x
 
### 78. subsets

1. 递归

### 169. majority-element

1. 哈希表：遍历记录每个元素出现的次数，大于len(nums)/2即为众数
2. 计数投票：用一个整数变量count做计数，遇到众数+1，非众数-1，最后count必定大于0，对应的candidate即为众数


### 17. letter-combinations-of-a-phone-number

1. 递归
