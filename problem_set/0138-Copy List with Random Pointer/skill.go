func copyRandomList(head *Node) *Node {
    m := make(map[*Node]*Node, 0)    // initiate an empty map
    dummy := new(Node)
    list, curr := dummy, head
    // Duplicate each node (without 'random' pointers) and put it into the map
    for curr != nil {
        n := &Node{
            Val: curr.Val,
        }
        list.Next = n
        m[curr] = n   // map old node to its new duplicate
        curr = curr.Next
        list = list.Next
    }
    list = dummy.Next
    // Assign 'random' pointers to duplicated nodes using the map
    for head != nil {
        list.Random = m[head.Random]
        list = list.Next
        head = head.Next
    }
    return dummy.Next  // return the head of the new list	
}

func copyRandomList(head *Node) *Node {
    // create duplicates and place them next to the originals
    cur := head 
    for cur != nil {
        node := &Node{
            Val: cur.Val,
            Next: cur.Next,
            Random: cur.Random,
        }
        cur.Next = node
        cur = node.Next
    }
    // fix 'random' pointers for duplicated nodes
    cur = head
    for cur != nil {
        if cur.Next.Random != nil {
            cur.Next.Random = cur.Next.Random.Next
        }
        cur = cur.Next.Next   // skip duplicate nodes
    }
    // separate the two lists
    dummy := new(Node)
    cur, origin := dummy, head
    for origin != nil {
        tmp := origin.Next
        origin.Next = tmp.Next
        origin = origin.Next
        cur.Next = tmp
        cur = cur.Next
    }
    return dummy.Next  // return the head of the new list 
}

