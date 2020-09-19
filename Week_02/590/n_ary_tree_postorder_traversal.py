from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        res = []

        def post_order(node: Node):
            if node is not None:
                for c in node.children:
                    post_order(c)
                res.append(node.val)

        post_order(root)
        return res
