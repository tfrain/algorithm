func detectCycle(head *ListNode) *ListNode {
    // return nil for empty list or single element list
    if head == nil || head.Next == nil {
        return nil
    }
    slow, fast := head.Next, head.Next.Next // initialize slow and fast pointers
    // loop through the list until a cycle is detected or end of list is reached
    for fast != nil && fast.Next != nil && fast != slow {
        fast = fast.Next.Next // move 2 steps
        slow = slow.Next // move 1 step
    }
    // return nil if no cycle is detected
    if fast == nil || fast.Next == nil {
        return nil
    }
    // to find the start of cycle, move slow back to head
    slow = head
    // then move slow and fast at the same pace until they meet
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    // return the node where slow and fast meet
    return slow
}