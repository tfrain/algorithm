package main

func sortList1(head *ListNode) *ListNode {
	// Iterate through each node in the list
	for i := head; i != nil; i = i.Next {
		// Compare and swap values with the next nodes
		for j := i.Next; j != nil; j = j.Next {
			if j.Val < i.Val {
				// Swap values if they are in the wrong order
				i.Val, j.Val = j.Val, i.Val
			}
		}
	}
	// Return the sorted list
	return head
}
