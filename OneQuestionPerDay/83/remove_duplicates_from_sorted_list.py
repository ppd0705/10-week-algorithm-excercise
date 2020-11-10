# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head is None:
            return head
        elments = {head.val}
        prev, cur = head, head.next
        while cur is not None:
            if cur.val not in elments:
                elments.add(cur.val)
                prev, cur = cur, cur.next
            else:
                prev.next = cur.next
                cur = cur.next
        return head
