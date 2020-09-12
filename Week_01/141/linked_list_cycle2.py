class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head is None or head.next is None:
            return False
        slow, fast = head, head.next
        while fast is not None and fast.next is not None:
            if slow == fast:
                return True
            slow = slow.next
            fast = fast.next.next
        return False

    def test(self):
        n1 = ListNode(1)
        n2 = ListNode(2)
        n3 = ListNode(3)
        n4 = ListNode(4)
        n2.next = n1
        n3.next = n2
        n4.next = n3

        assert not self.hasCycle(n1)
        assert not self.hasCycle(n2)
        assert not self.hasCycle(n3)
        assert not self.hasCycle(n4)
        n1.next = n3
        assert self.hasCycle(n1)
        assert self.hasCycle(n2)
        assert self.hasCycle(n3)
        assert self.hasCycle(n4)

        print("all done")


if __name__ == "__main__":
    Solution().test()
