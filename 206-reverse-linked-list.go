func reverseList(head *ListNode) *ListNode {
    var previous *ListNode = nil
    current := head
    for current != nil {
        next := current.Next
        current.Next = previous
        previous = current
        current = next 
    }
    return previous
}
