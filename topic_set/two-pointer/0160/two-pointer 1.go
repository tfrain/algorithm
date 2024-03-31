func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // Initialize two pointers p1 and p2 at the head of A and B.
    p1, p2 := headA, headB

    // Loop until the two pointers meet.
    for p1 != p2 {
        // If p1 reaches the end of list A, redirect it to the head of list B.
        if p1 == nil {
            p1 = headB
            // Skip the next step since we have already moved p1.
            continue
        }
        // If p2 reaches the end of list B, redirect it to the head of list A.
        if p2 == nil {
            p2 = headA
            // Skip the next step since we have already moved p2.
            continue
        }
        // Move p1 and p2 one step forward.
        p1 = p1.Next
        p2 = p2.Next
    }
    // Return the intersection node. If there is no intersection, p1 will be nil.
    return p1
}