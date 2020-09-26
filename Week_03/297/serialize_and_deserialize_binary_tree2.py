class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if root is None:
            return "#"
        return str(root.val) + "," + self.serialize(root.left) + "," + self.serialize(root.right)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        vals = iter(data.split(","))

        def dfs():
            v = next(vals)
            if v == "#":
                return None
            node = TreeNode(int(v))
            node.left = dfs()
            node.right = dfs()
            return node

        return dfs()

    def test(self):

        node1 = TreeNode(1)
        node2 = TreeNode(2)
        node3 = TreeNode(3)
        node4 = TreeNode(4)
        node5 = TreeNode(5)

        node1.left = node2
        node1.right = node3
        node2.left = node4
        node3.left = node5

        data = self.serialize(node1)
        print(data)
        node = self.deserialize(data)
        assert node.val == 1
        print("well done")


if __name__ == "__main__":
    Codec().test()
