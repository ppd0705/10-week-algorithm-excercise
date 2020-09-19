from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        res = []

        def pre_order(node: TreeNode):
            if node is not None:
                res.append(node.val)
                pre_order(node.left)
                pre_order(node.right)

        pre_order(root)
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
