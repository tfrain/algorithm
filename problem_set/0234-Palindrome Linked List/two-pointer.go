// isPalindrome checks if the LinkedList is a Palindrome
func isPalindrome(head *ListNode) bool {
    // prev is a pointer to store the previous Node
    var prev *ListNode

    // Initializing slow and fast pointers to traverse the linked list
    slow, fast := head, head

    // Loop through the linked list
    // Reverse the first half of the list while finding the middle
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next // move 2 steps

        // save the next node before overriding the link
        nxt := slow.Next 

        // Reverse the current node
        slow.Next = prev 

        // Move one step ahead
        prev = slow 
        slow = nxt 
    }
    // Handling the case when the linked list has odd number of nodes.
    // Fast wasn't nil so we skip the central node
    if fast != nil {
        slow = slow.Next
    }

    // Loop for comparison
    for slow != nil {
        // If values at the current nodes in the reversed first half
        // and the second half aren't the same, it's not a palindrome
        if slow.Val != prev.Val {
            return false
        }
        
        // Move one step ahead for comparison in the next iteration
        slow = slow.Next
        prev = prev.Next
    }
    // If all the nodes matched, it's a palindrome
    return true
}