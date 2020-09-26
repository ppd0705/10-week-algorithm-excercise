class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if root is None:
            return 0
        dep = 0

        queue = [root]
        while queue:
            n = len(queue)
            dep += 1
            for i in range(n):
                if queue[i].left is None and queue[i].right is None:
                    return dep
                if queue[i].left is not None:
                    queue.append(queue[i].left)

                if queue[i].right is not None:
                    queue.append(queue[i].right)
            queue = queue[n:]
        return -1
