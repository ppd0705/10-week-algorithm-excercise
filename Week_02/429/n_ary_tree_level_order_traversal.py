from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        res = []

        def level_order(node: Node, level: int):
            if node is not None:
                if len(res) == level:
                    res.append([node.val])
                else:
                    res[level].append(node.val)
                for c in node.children:
                    level_order(c, level + 1)

        level_order(root, 0)
        return res
