from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        res = []

        stack = [root]
        while len(stack) > 0:
            node = stack.pop()
            if node is not None:
                res.append(node.val)
                for c in node.children:
                    stack.append(c)
        return res[::-1]
