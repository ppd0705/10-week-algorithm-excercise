from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        res = []

        stack = [root]
        while stack:
            node = stack.pop()
            if node is not None:
                res.append(node.val)
                for i in range(len(node.children) - 1, -1, -1):
                    stack.append(node.children[i])
        return res
