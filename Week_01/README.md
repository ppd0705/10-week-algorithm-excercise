学习笔记


|#|Title|Solutions|
|---|---|------|
|70|[climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs) | 递推([Go](70/climbing_stairs.go),[Py](70/climbing_stairs.py))|
|11|[container-with-most-water](https://leetcode-cn.com/problems/container-with-most-water) | 双指针([Go](11/container_with_most_water.go),[Py](11/container_with_most_water.py))|
|283|[move-zeroes](https://leetcode-cn.com/problems/move-zeroes) | 统计0的个数([Go](283/move_zeros.go),[Py](283/move_zeros.py)), 快慢指针([Go](283/move_zeros.go))|



## 题解

### 11. Container With Most Water

1. 暴力法：两层for循环
2. 双指针：缩减搜索空间
   对于i,j(i < j),如果h[i] <= h[j],那么任意的j1(i < j1 < j),必定会有h[i] * (j1 -i) < h[i] * (j-i)
   
   
### 283. move-zeroes

1. 统计0的个数，遇到非0时 nums[i-count] = nums[i] 
2. 慢指针指向非0元素