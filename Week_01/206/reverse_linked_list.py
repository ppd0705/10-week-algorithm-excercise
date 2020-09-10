class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None:
            return

        head_new = head
        while head.next is not None:
            next_ = head.next
            next_.next, head.next = head_new, next_.next
            head_new = next_
        return head_new

    def test(self):
        node1 = ListNode(1)
        node2 = ListNode(2)
        node3 = ListNode(3)
        node4 = ListNode(4)

        node2.next = node1
        node3.next = node2
        node4.next = node3

        ans = self.reverseList(node1)
        assert ans == node1

        ans = self.reverseList(node2)
        assert ans == node1

        ans = self.reverseList(node1)
        assert ans == node2

        ans = self.reverseList(node4)
        assert ans == node1

        print("success")


if __name__ == "__main__":
    Solution().test()
