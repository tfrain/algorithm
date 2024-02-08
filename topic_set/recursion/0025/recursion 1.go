//reverseKGroup reverses k nodes in a linked list
func reverseKGroup(head *ListNode, k int) *ListNode {
    // check if head is nil, in that case just return nil
    if head == nil {
        return nil
    }
    first := head
    for i := 1; i < k; i++ {
        first = first.Next
        // if there are less than k nodes in the list, just return the list as it is.
        if first == nil {
            return head
        }
    }
    // other nodes in the list
    others := first.Next
    first.Next = nil
    // reverse the k nodes
    newHead := reverse(head)
    // recursively reverse the remaining nodes
    head.Next = reverseKGroup(others, k)
    return newHead
}

//reverse reverses nodes in a linked list
func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    // main loop to reverse the list nodes
    for head != nil {
        nxt := head.Next
        // reverse the 'next' pointer
        head.Next = prev
        prev = head
        head = nxt
    }
    // return the new head after reversing
    return prev
}