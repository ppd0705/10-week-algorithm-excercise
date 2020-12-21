学习笔记


|#|Title|Solutions|
|---|---|------|
|990|[satisfiability-of-equality-equations](https://leetcode-cn.com/problems/satisfiability-of-equality-equations) | 并查集([Go](./990/satisfiability_of_equality_equations.go)）|
|721|[accounts-merge](https://leetcode-cn.com/problems/accounts-merge) | 并查集([Go](./721/accounts_merge.go)）|


## 题解

### 721. Accounts Merge

1. 并查集：
    使用一个字典记录email对应账户的index,如果已经在字典里面出现了，就说明需要合并用户