学习笔记


|#|Title|Solutions|
|---|---|------|
|191|[number-of-1-bits](https://leetcode-cn.com/problems/number-of-1-bits) | 依次除2([Go](../Week_08/191/number_of_1_bits.go),[Py](../Week_08/191/number_of_1_bits.py)),按位与n-1([Go](../Week_08/191/number_of_1_bits2.go),[Py](../Week_08/191/number_of_1_bits2.py))|
|190|[reverse_bits](https://leetcode-cn.com/problems/reverse_bits) | 迭代([Go](../Week_08/190/reverse_bits.go),[Py](../Week_08/190/reverse_bits.py)),分治位运算([Go](../Week_08/190/reverse_bits2.go),[Py](../Week_08/190/reverse_bits2.py))|
|338|[counting-bits](https://leetcode-cn.com/problems/counting-bits) | 动态规划1([Go](../Week_08/338/counting_bits.go),[Py](../Week_08/338/counting_bits.py)),动态规划2([Go](../Week_08/338/counting_bits2.go),[Py](../Week_08/338/counting_bits2.py))|
|231|[power-of-two](https://leetcode-cn.com/problems/power-of-two) | 位运算1([Go](../Week_08/231/power_of_two.go),[Py](../Week_08/231/power_of_two.py)),位运算2([Go](../Week_08/231/power_of_two2.go),[Py](../Week_08/231/power_of_two2.py))|
|326|[power-of-three](https://leetcode-cn.com/problems/power-of-three) | 累除法([Go](../Week_08/326/power_of_three.go),[Py](../Week_08/326/power_of_three.py)),累乘法([Go](../Week_08/326/power_of_three2.go),[Py](../Week_08/326/power_of_three2.py))|
|342|[power-of-four](https://leetcode-cn.com/problems/power-of-four) | 累除法([Go](../Week_08/342/power_of_four.go),[Py](../Week_08/342/power_of_four.py)),位运算([Go](../Week_08/342/power_of_four2.go),[Py](../Week_08/342/power_of_four2.py))|
|146|[lru-cache](https://leetcode-cn.com/problems/lru-cache) | 双向链表+字典([Go](../Week_08/146/lru_cache.go),[Py](../Week_08/146/lru_cache.py))|
|-1|[sort](sort) | 冒泡/选择/插入/归并/快速/归并排序([Go](../Week_08/sort/sort.go),[Py](../Week_08/sort/sort.py))|




## 题解

### 191. number-of-1-bits

1. 除2：依次除2，看余数是否是1 
2. 按位与n-1: n & n-1 会消去最后一个1

### 190. reverse-bits

1. 除2：
2. 分治：依次以16、8、4、2、1两两交换

### 338. counting-bits

1. 动态规划1： dp[i] = dp[i//2] + i % 2
2. 动态规划2： dp[i] = dp[i & (i-1)] + 1

### 231. power-of-two

1. 位运算1：  n & (n-1)为消去最后一个1
2. 位运算2： n & -n会保留最后一个1，其他位都置0

### 326. power-of-three

1. 累除法
2. 累乘法

### 342. power-of-four

1. 累除法
2. 位运算

### 146. lru_cache

1. 双向链表+字典