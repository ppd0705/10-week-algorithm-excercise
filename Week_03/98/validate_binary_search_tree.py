# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        return self.dfs(root, -2 ** 63, 2 ** 64 - 1)

    def dfs(self, node: TreeNode, left: int, right: int) -> bool:
        if node is None:
            return True

        return left < node.val < right and self.dfs(node.left, left, node.val) and self.dfs(node.right, node.val, right)
