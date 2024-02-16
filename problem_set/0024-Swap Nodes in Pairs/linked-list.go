func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head.Next
    for head != nil && head.Next != nil {
        first, second := head, head.Next
        nxt := second.Next
        if nxt != nil && nxt.Next != nil {
            nxt = nxt.Next
        }
        head = second.Next
        first.Next, second.Next = nxt, first
    }
    return newHead
}