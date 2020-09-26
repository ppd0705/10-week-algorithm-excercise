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
        vals = []
        queue = [root]
        while queue:
            node = queue.pop(0)
            if node is None:
                vals.append("#")
            else:
                vals.append(str(node.val))
                queue.append(node.left)
                queue.append(node.right)
        return ",".join(vals)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        vals = data.split(",")
        if vals[0] == "#":
            return None

        idx = 0
        root = TreeNode(int(vals[idx]))
        queue = [root]

        while queue:
            node = queue.pop(0)
            idx += 1
            if vals[idx] != "#":
                left = TreeNode(int(vals[idx]))
                node.left = left
                queue.append(left)
            idx += 1
            if vals[idx] != "#":
                right = TreeNode(int(vals[idx]))
                node.right = right
                queue.append(right)
        return root

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
