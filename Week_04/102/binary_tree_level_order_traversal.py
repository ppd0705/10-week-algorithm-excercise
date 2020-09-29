from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []

        if root is None:
            return res

        queue = [root]
        level = 0
        while queue:
            n = len(queue)
            for i in range(n):
                if level == len(res):
                    res.append([queue[i].val])
                else:
                    res[level].append(queue[i].val)
                if queue[i].left is not None:
                    queue.append(queue[i].left)
                if queue[i].right is not None:
                    queue.append(queue[i].right)
            level += 1
            queue = queue[n:]
        return res
