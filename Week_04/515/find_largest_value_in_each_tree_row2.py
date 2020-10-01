from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        res = []

        def dfs(node: True, level: int):
            if node is None:
                return
            if level == len(res):
                res.append(node.val)
            elif node.val > res[level]:
                res[level] = node.val

            dfs(node.left, level + 1)
            dfs(node.right, level + 1)

        dfs(root, 0)
        return res
