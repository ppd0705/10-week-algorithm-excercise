学习笔记

|#|Title|Solutions|
|---|---|------|
|191|[number-of-1-bits](https://leetcode-cn.com/problems/number-of-1-bits) | 依次除2([Go](../Week_08/191/number_of_1_bits.go),[Py](../Week_08/191/number_of_1_bits.py)),按位与n-1([Go](../Week_08/191/number_of_1_bits2.go),[Py](../Week_08/191/number_of_1_bits2.py))|
|190|[reverse_bits](https://leetcode-cn.com/problems/reverse_bits) | 迭代([Go](../Week_08/190/reverse_bits.go),[Py](../Week_08/190/reverse_bits.py)),分治位运算([Go](../Week_08/190/reverse_bits2.go),[Py](../Week_08/190/reverse_bits2.py))|
|338|[counting-bits](https://leetcode-cn.com/problems/counting-bits) | 动态规划1([Go](../Week_08/338/counting_bits.go),[Py](../Week_08/338/counting_bits.py)),动态规划2([Go](../Week_08/338/counting_bits2.go),[Py](../Week_08/338/counting_bits2.py))|


## 题解

### 191. number-of-1-bits

1. 除2：依次除2，看余数是否是1 
2. 按位与n-1: n & n-1 为消去最后一位1

### 190. reverse-bits

1. 除2：
2. 分治：依次以16、8、4、2、1两两交换

### 338. counting-bits

1. 动态规划1： dp[i] = dp[i//2] + i % 2
2. 动态规划2： dp[i] = dp[i & (i-1)] + 1