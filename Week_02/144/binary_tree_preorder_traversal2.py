from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        res = []
        stack = [root]

        while len(stack) > 0:
            node = stack.pop()
            if node is not None:
                res.append(node.val)
                if node.right is not None:
                    stack.append(node.right)
                if node.left is not None:
                    stack.append(node.left)
        return res

    def test(self):
        n1 = TreeNode(1)
        n2 = TreeNode(2)
        n3 = TreeNode(3)

        n1.right = n2
        n2.left = n3

        ans = self.preorderTraversal(n1)
        assert ans == [1, 2, 3], f"ans: {ans}"
        print("well done")


if __name__ == "__main__":
    Solution().test()
