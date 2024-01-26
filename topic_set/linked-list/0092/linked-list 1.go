func reverseBetween(head *ListNode, m int, n int) *ListNode {
    // Initializing the dummy node with the head of the linked list
    dummy := &ListNode{
        Next: head,
    }
    // Initializing a pointer to traverse the linked list till we reach 'm'
    list := dummy
    for i := 1; i < m; i++ {
        list = list.Next
    }
    // Setting up pointers for list reversal: 'prev' points to the m'th node and 'curr' points to the (m+1)'th node
    prev, curr := list.Next, list.Next.Next
    // Starting the reversing process by changing the link for each node in the sublist from 'curr' to 'prev'
    for i := m; i < n; i++ {
        nxt := curr.Next   // Store the next node
        curr.Next = prev   // Reverse the link
        prev = curr        // Update 'prev' for next reversal
        curr = nxt         // Move 'curr' to the next node
    }
    // After reversal, connect the reversed sublist back with the rest of the list
    list.Next.Next = curr // Connect the originally m'th node with the (n+1)'th node
    list.Next = prev      // Connect the (m-1)'th node with the new head of reversed sublist (originally n'th node)
    // 'dummy.Next' is the new head of the entire modified linked list
    return dummy.Next
}