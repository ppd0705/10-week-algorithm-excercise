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
|21|[merge-two-sorted-lists](https://leetcode-cn.com/problems/merge-two-sorted-lists) | 迭代([Go](21/merge_two_sorted_lists.go), [Py](21/merge_two_sorted_lists.py)), 递归([Go](21/merge_two_sorted_lists2.go),[Py](21/merge_two_sorted_lists2.py))|
|141|[linked-list-cycle](https://leetcode-cn.com/problems/linked-list-cycle) | 哈希([Go](141/linked_list_cycle.go), [Py](141/linked_list_cycle.py)), 快慢指针([Go](141/linked_list_cycle2.go),[Py](141/linked_list_cycle2.py)), 翻转链表([Go](141/linked_list_cycle3.go))|
|142|[linked-list-cycle-ii](https://leetcode-cn.com/problems/linked-list-cycle-ii) | 快慢指针([Go](142/linked_list_cycle_ii.go), [Py](142/linked_list_cycle_ii.py))|
|25|[reverse-nodes-in-k-group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) | 快慢指针([Go](25/reverse_nodes_in_k_group.go), [Py](25/reverse_nodes_in_k_group.py))|
|20|[valid-parentheses](https://leetcode-cn.com/problems/valid-parentheses) | 栈([Go](20/valid_parentheses.go), [Py](20/valid_parentheses.py))|
|155|[min-stack](https://leetcode-cn.com/problems/min-stack) | 双栈([Go](155/min_stack.go), [Py](155/min_stack.py))|
|84|[largest-rectangle-in-histogram](https://leetcode-cn.com/problems/largest-rectangle-in-histogram) | 暴力法([Go](84/largest_rectangle_in_histogram.go)),暴力法优化([Go](84/largest_rectangle_in_histogram2.go)),单调栈([Go](84/largest_rectangle_in_histogram3.go), [Py](84/largest_rectangle_in_histogram3.py))|
|239|[sliding-window-maximum](https://leetcode-cn.com/problems/sliding-window-maximum) | 单调递减队列([Go](239/sliding_window_maximum.go), [Py](239/sliding_window_maximum.py))|
|641|[design-circular-deque](https://leetcode-cn.com/problems/design-circular-deque) | 双指针([Go](641/design_circular_deque.go)),双指针优化([Go](641/design_circular_deque2.go),[Py](641/design_circular_deque2.py))|
|42|[trapping-rain-water](https://leetcode-cn.com/problems/trapping-rain-water) | 单调栈([Go](42/trapping_rain_water.go),[Py](42/trapping_rain_water.py))|
|26|[remove-duplicates-from-sorted-array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) | 快慢指针([Go](26/remove_duplicates_from_sorted_array.go),[Py](26/remove_duplicates_from_sorted_array.py))|
|189|[rotate-array](https://leetcode-cn.com/problems/rotate-array) | 暴力法([Go](189/rotate_array.go)), 三次翻转([Go](189/rotate_array2.go),[Py](189/rotate_array2.py))|
|88|[merge-sorted-array](https://leetcode-cn.com/problems/merge-sorted-array) | 暴力法([Go](88/merge_sorted_array.go)), 双指针从后向前([Go](88/merge_sorted_array2.go),[Py](88/merge_sorted_array2.py))|



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

### 21. merge-two-sorted-lists

1. 迭代法: 构造一个dummy node 当前置节点
2. 递归法


### 141. linked-list-cycle 

1. 哈希表法: 用哈希表记录是否访问过
2. 双指针法: 慢指针一次走一步，快指针一次走两步，如果没环，快指针会先指向nil
3. 翻转链表法: 每走一步，将Next之前前置节点，如果有环，会重新回到head节点

### 142. linked-list-cycle-ii 

快慢指针法： 
  - 假设环入口长度为a,环长度为b;则走a+ n * b步时在环入口处
  - 第一次相遇时慢指针走的不步数s=nb(f == 2 * s == s + nb);
  - 慢指指针再走a步即可到达环入口，让快指针按步长为1从head和slow同步移动，则再次相遇时即走了a步
  
  
### 25. reverse-nodes-in-k-group

快慢指针法：
  - 快指针先试探能否走k步
  - 若能依次翻转
  - 从快指针的next继续下一轮


### 20. valid-parentheses

哈希表存映射，栈存左括号

### 155. min stack

双栈法：给最小值也建个栈

### 84. largest-rectangle-in-histogram

1. 暴力法: 以每个bar为高度基准，找到左右边界
2. 暴力法优化: 左边界在height[i] <= height[i-1]时 bounder[i] = bounder[i-1],右边界同理
3. 单调栈: 栈存每个bar的索引，如果新的bar height < heights[stack[-1]]时需要出栈


### 239. sliding-window-maximum
队列法：
  - 移除旧窗口的元素
  - 移除当前窗口比新元素小的
  - 新元素入队
  - 更新最大窗口集
  
### 641. design-circular-dequeue

1. 双指针：建立大小为k的数据，头尾指针都指向当前插入数据的位置，并用len记录当前长度以判断是否空或满
2. 双指针优化: 为了避免头尾指针在空队列时候的耦合，让头指针指向头元素

### 42. trapping rain water

单调栈: 栈存每个bar的索引，如果新的bar height > heights[stack[-1]]时需要出栈


### 26. remove-duplicates-from-sorted-array

快慢指针: 快指针遍历，慢指针记录当前最新不重复的数的位置，关键是如何写的优雅


### 189. rotate-array

1. 暴力法: 每次移动最后一个元素到队首，移动k次
2. 三次翻转法: 先全部翻转，再翻转[0,k-1], [k, len(nums)-1] 

### 88. merge-sorted-array

1. 暴力法: 逐个加入nums2, 插入时整体copy nums1后移
2. 双指针法:从后向前比较num1和nums2