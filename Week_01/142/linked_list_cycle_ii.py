class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        fast = slow = head
        while True:
            if fast is None or fast.next is None:
                return None
            fast = fast.next.next
            slow = slow.next
            if fast is slow:
                break
        fast = head

        while fast is not slow:
            fast = fast.next
            slow = slow.next
        return fast
