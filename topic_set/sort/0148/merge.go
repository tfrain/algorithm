package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge Sort Implementation
func sortList(head *ListNode) *ListNode {
	// Base case: if the list is empty or has only one node
	if head == nil || head.Next == nil {
		return head
	}
	// Use slow and fast pointers to find the middle of the list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// Divide the list into two halves
	firstTail := slow
	slow = slow.Next
	firstTail.Next = nil // Cut the list at the middle

	// Recursively sort both halves
	first, second := sortList(head), sortList(slow)

	// Merge the two sorted halves
	return merge(first, second)
}

func merge(first, second *ListNode) *ListNode {
	// Dummy node to start the merged list
	dummy := new(ListNode)
	list := dummy
	// Merge nodes from first and second lists based on their values
	for first != nil && second != nil {
		if first.Val < second.Val {
			list.Next = first
			first = first.Next
		} else {
			list.Next = second
			second = second.Next
		}
		list = list.Next
	}
	// Attach any remaining nodes from either list
	if first != nil {
		list.Next = first
	}
	if second != nil {
		list.Next = second
	}
	// Return the merged list, starting from the next of dummy
	return dummy.Next
}
