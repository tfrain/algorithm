package main

// insertion sort algorithm
func sortList2(head *ListNode) *ListNode {
	// Return nil if the list is empty
	if head == nil {
		return nil
	}

	// Dummy node to help with insertion
	dummy := &ListNode{
		Next: head,
	}

	// Traverse the list to sort it
	for prev, curr := head, head.Next; curr != nil; curr = curr.Next {
		// If the current value is in order, move to the next node
		if prev.Val <= curr.Val {
			prev = curr
			continue
		}

		// Find the correct position for 'curr'
		j := dummy
		for j.Next.Val < curr.Val {
			j = j.Next
		}

		// Re-link nodes to insert 'curr' at the correct position
		prev.Next = curr.Next
		curr.Next = j.Next
		j.Next = curr

		// Reset 'curr' to the next node to continue sorting
		curr = prev
	}

	// Return the sorted list, excluding the dummy node
	return dummy.Next
}
