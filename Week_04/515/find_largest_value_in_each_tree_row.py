from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        res = []
        if root == None:
            return res

        queue = [root]
        while queue:
            n = len(queue)
            val = queue[0].val
            for i in range(n):
                if queue[i].val > val:
                    val = queue[i].val
                if queue[i].left is not None:
                    queue.append(queue[i].left)
                if queue[i].right is not None:
                    queue.append(queue[i].right)
            res.append(val)
            queue = queue[n:]
        return res
