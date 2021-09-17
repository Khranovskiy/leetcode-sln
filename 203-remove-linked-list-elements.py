class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        result = ListNode(0)
        result.next = head
        
        current = result
        while current.next:

            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        
        return result.next
