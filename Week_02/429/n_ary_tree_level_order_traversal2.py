from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        res = []
        level = 0
        stack = [root]
        while stack:
            num = len(stack)
            for i in range(num):
                if stack[i] is not None:
                    if len(res) == level:
                        res.append([stack[i].val])
                    else:
                        res[level].append(stack[i].val)
                    for c in stack[i].children:
                        stack.append(c)
            stack = stack[num:]
            level += 1
        return res
