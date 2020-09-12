class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        dummpy = ListNode(0)
        dummpy.next = head
        prev = dummpy
        fast = slow = head

        while True:
            cnt = 1

            while cnt < k and fast is not None:
                fast = fast.next
                cnt += 1

            if fast is None:
                break

            prev.next = fast
            prev = slow
            next_ = slow.next
            prev.next = fast.next

            while next_ is not prev.next:
                next_.next, next_, slow = slow, next_.next, next_

            fast = slow = prev.next

        return dummpy.next

    def test(self):
        n1 = ListNode(1)
        n2 = ListNode(2)
        n3 = ListNode(3)
        n4 = ListNode(4)
        n5 = ListNode(5)
        n1.next = n2
        n2.next = n3
        n3.next = n4
        n4.next = n5

        k = 3
        head = self.reverseKGroup(n1, k)

        while head is not None:
            print(f"{head.val}-->", end="")
            head = head.next


if __name__ == "__main__":
    Solution().test()
