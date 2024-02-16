func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	} // Head node that'll point to the resulting linked list
	p := dummy // Pointer that'll traverse and modify ret
	carry := 0 // Variable to handle digit carry over during addition
	for l1 != nil || l2 != nil || carry > 0 {
		val := carry // store's current digit's value
		if l1 != nil {
			val += l1.Val // Add l1 node's value if it exists
			l1 = l1.Next  // Move to next node in l1
		}
		if l2 != nil {
			val += l2.Val // Add l2 node's value if it exists
			l2 = l2.Next  // Move to next node in l2
		}
		carry = val / 10             // Update carry; we obtain it by integer division of sum by 10
		val = val % 10               // Update val to only store a single digit
		p.Next = &ListNode{Val: val} // Add new node with our current sum
		p = p.Next                   // Move our modifying pointer to the last node in resultant list
	}
	return dummy.Next // Return the head of resultant list, as ret points to an unnecessary node
}