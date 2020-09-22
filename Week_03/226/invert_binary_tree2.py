class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        queue = [root]
        while queue:
            node = queue.pop()
            if node:
                node.right, node.left = node.left, node.right
                queue.append(node.right)
                queue.append(node.left)
        return root
