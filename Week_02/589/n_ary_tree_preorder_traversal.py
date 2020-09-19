from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        res = []

        def pre_order(node: Node):
            if node is not None:
                res.append(node.val)
                for c in node.children:
                    pre_order(c)

        pre_order(root)
        return res
