from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        res = []

        while head is not None:
            res.append(head.val)
            head = head.next
        return res[::-1]