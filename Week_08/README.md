学习笔记

|#|Title|Solutions|
|---|---|------|
|191|[number-of-1-bits](https://leetcode-cn.com/problems/number-of-1-bits) | 依次除2([Go](../Week_08/191/number_of_1_bits.go),[Py](../Week_08/191/number_of_1_bits.py)),按位与n-1([Go](../Week_08/191/number_of_1_bits2.go),[Py](../Week_08/191/number_of_1_bits2.py))|


## 题解

### 191. number-of-1-bits

1. 除2：依次除2，看余数是否是1 
2. 按位与n-1: n & n-1 为消去最后一位1