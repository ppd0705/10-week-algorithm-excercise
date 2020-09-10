学习笔记


|#|Title|Solutions|
|---|---|------|
|70|[climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs) | 递推([Go](70/climbing_stairs.go),[Py](70/climbing_stairs.py))|
|11|[container-with-most-water](https://leetcode-cn.com/problems/container-with-most-water) | 双指针([Go](11/container_with_most_water.go),[Py](11/container_with_most_water.py))|
|283|[move-zeroes](https://leetcode-cn.com/problems/move-zeroes)| 统计0的个数([Go](283/move_zeros.go),[Py](283/move_zeros.py)), 快慢指针([Go](283/move_zeros.go))|
|1|[two-sum](https://leetcode-cn.com/problems/two-sum)| 先排序后双指针夹逼([Go](1/two_sum.go)), 字典缓存差值的索引([Go](1/two_sum_2.go),[Py](1/two_sum.py))|
|15|[3sum](https://leetcode-cn.com/problems/3sum)| 先排序后双指针夹逼([Go](15/3sum_2.go),[Py](15/3sum.py)), 字典缓存差值的索引([Go](15/3sum.go))|
|206|[reverse-linked-list](https://leetcode-cn.com/problems/reverse-linked-list)| 迭代([Go](206/reverse_linked_list.go),[Py](206/reverse_linked_list.py)), 递归([Go](206/reverse_linked_list2.go),[Py](206/reverse_linked_list2.py))|
|24|[swap-nodes-in-pairs](https://leetcode-cn.com/problems/swap-nodes-in-pairs)| 迭代([Go](24/swap_nodes_in_pairs2.go),[Py](24/swap_nodes_in_pairs2.py)), 递归([Go](24/swap_nodes_in_pairs.go),[Py](24/swap_nodes_in_pairs.py))|




## 题解

### 11. Container With Most Water

1. 暴力法：两层for循环
2. 双指针：缩减搜索空间
   对于i,j(i < j),如果h[i] <= h[j],那么任意的j1(i < j1 < j),必定会有h[i] * (j1 -i) < h[i] * (j-i)
   
   
### 283. move-zeroes

1. 统计0的个数，遇到非0时 nums[i-count] = nums[i] 
2. 慢指针指向非0元素

### 1. two-sum

计算差值是否在字典里，如果是则已找到，否则将当前值存入字典

### 15. 3sum

先排序，然后快慢指指针夹逼，另外要注意去重


### 206. reverse linked list

1. 迭代法：使用一个指正指向最新的head
2. 递归法：比较难以理解，`head.next.next = head; head.next = None`

### 24. swap-nodes-in-pairs

1. 迭代法：构造一个dummy node 当前置节点
2. 递归法