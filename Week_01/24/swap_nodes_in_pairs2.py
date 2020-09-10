class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __repr__(self):
        return f"[Node]{self.val}"


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        dummy = ListNode(0)
        dummy.next = head
        prev = dummy

        while head is not None and head.next is not None:
            prev.next = head.next
            prev = head
            head.next.next, head.next = head, head.next.next
            head = head.next
        return dummy.next

    def test(self):
        node1 = ListNode(1)
        node2 = ListNode(2)
        node3 = ListNode(3)
        node4 = ListNode(4)

        node2.next = node1
        node3.next = node2
        node4.next = node3

        ans = self.swapPairs(node1)
        assert ans == node1

        ans = self.swapPairs(node2)
        assert ans == node1
        #
        ans = self.swapPairs(node1)
        assert ans == node2

        ans = self.swapPairs(node4)
        assert ans == node3

        print("success")


if __name__ == "__main__":
    Solution().test()


