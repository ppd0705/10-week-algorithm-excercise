from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        res = []

        while root is not None:
            res.append(root.val)
            if root.left is None:
                root = root.right
            else:
                pre = root.left
                while pre.right is not None:
                    pre = pre.right
                pre.right = root.right
                root = root.left

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
