func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Define a dummy node and link it to the head
    dummy := new(ListNode)
    dummy.Next = head
    
    // Initialize slow and fast pointers at the dummy node
    slow, fast := dummy, dummy
    
    // Move the fast pointer n nodes ahead
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    
    // Move both pointers step by step until fast reaches the end
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    
    // Delete the nth node from the end
    slow.Next = slow.Next.Next
    
    // Return the modified list
    return dummy.Next
}