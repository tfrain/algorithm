// func mergeKLists receives an array of sorted linked lists and merge them into one sorted linked list
func mergeKLists(lists []*ListNode) *ListNode {
    // If there are no lists, return nil
    if len(lists) == 0 {
        return nil
    }
    // If there is only one list, return that list
    if len(lists) == 1 {
        return lists[0]
    }
    // For more than one list, we start merging them in pairs
    for len(lists) > 1 {
        list1, list2 := lists[0], lists[1]
        merged := mergeList(list1, list2) 
        lists = append(lists, merged) // Append the merged list to the end of the array
        // Remove the two lists we just processed
        lists = lists[2:] 
    }
    return lists[0] // Return the final merged list
}

// func mergeList receives two lists, and merge them into one sorted linked list
func mergeList(listA, listB *ListNode) *ListNode {
    dummy := new(ListNode) // Create a dummy node to hold the merged list
    list := dummy // Pointer to the last added node in the merged list
    // Compare the heads of the two lists and add the smaller one to the merged list
    for listA != nil && listB != nil {
        if listA.Val < listB.Val {
            list.Next = listA
            listA = listA.Next
        } else {
            list.Next = listB
            listB = listB.Next
        }
        list = list.Next // Move the merge list pointer forward
    }

    // If there are nodes left in either of the lists, append them to the end of the merged list
    if listA != nil {
        list.Next = listA
    }
    if listB != nil {
        list.Next = listB
    }
    // Using the Next attribute of the dummy node will skip the initial dummy node, giving us our new merged list
    return dummy.Next 
}