class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if root is None:
            return 0
        level = 0
        queue = [root]

        while queue:
            n = len(queue)
            level += 1

            for i in range(n):
                if queue[i].left is not None:
                    queue.append(queue[i].left)
                if queue[i].right is not None:
                    queue.append(queue[i].right)

            queue = queue[n:]
        return level
