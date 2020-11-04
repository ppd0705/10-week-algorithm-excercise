class DNode:
    __slots__ = ("key", "value", "prev", "next")


class LRUCache(object):
    def __init__(self, capacity):
        self.cap = capacity
        self.root = DNode()
        self.root.next, self.root.prev = self.root, self.root
        self.root.key = self.root.value = -1
        self.cache = {}

    def move_to_end(self, node: DNode):
        node.next.prev, node.prev.next = node.prev, node.next
        node.prev, node.next = self.root.prev, self.root
        self.root.prev.next, self.root.prev = node, node

    def get(self, key: int):
        if key not in self.cache:
            return -1
        node = self.cache[key]
        self.move_to_end(node)
        return node.value

    def put(self, key: int, value: int):
        if key in self.cache:
            self.cache[key].value = value
            self.move_to_end(self.cache[key])
        else:
            if len(self.cache) < self.cap:
                node = DNode()
                node.key, node.value = key, value
                node.prev, node.next = self.root.prev, self.root
                self.root.prev.next, self.root.prev = node, node
                self.cache[key] = node
            else:
                self.root.key = key
                self.root.value = value
                self.cache[key] = self.root
                self.cache.pop(self.root.next.key)
                self.root = self.root.next
