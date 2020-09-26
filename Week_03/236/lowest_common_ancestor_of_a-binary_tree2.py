class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        parent = {root: None}
        visited = {}

        def dfs(node: TreeNode):
            if node is not None:
                if node.left is not None:
                    parent[node.left] = node
                    dfs(node.left)
                if node.right is not None:
                    parent[node.right] = node
                    dfs(node.right)

        dfs(root)
        while p is not None:
            visited[p] = True
            p = parent[p]

        while q is not None:
            if q in visited:
                return q
            q = parent[q]

        return None

    def test(self):

        node1 = TreeNode(1)
        node2 = TreeNode(2)
        node3 = TreeNode(3)
        node4 = TreeNode(4)

        node1.left = node2
        node1.right = node3
        node2.left = node4

        assert self.lowestCommonAncestor(node1, node2, node3) is node1

        print("well done")


if __name__ == "__main__":
    Solution().test()
