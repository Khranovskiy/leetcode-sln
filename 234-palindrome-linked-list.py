# [ 1  2  2  1]
# [ 1  2  
#         1  2->nil]
        
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if not head:
            return None
        
        middle = self.middleNode(head)
        reverse = self.reverseList(middle)

        while reverse:
            if head.val != reverse.val:
                return False

            reverse = reverse.next
            head = head.next

        return True
    
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        current = head
        while current:
            next = current.next
            current.next = prev
            prev = current
            current = next
        return prev
    
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fast = slow = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next         
        return slow



