// func mergeKLists receives an array of sorted linked lists and merge them into one sorted linked list
func mergeKLists(lists []*ListNode) *ListNode {
    pq := make(PQ, 0)
    for _, list := range lists {
        if list != nil {
            pq = append(pq, list)
        }
    }
    // If there are no lists, return nil
    if len(pq) == 0 {
        return nil
    }
    heap.Init(&pq) // Initialize our heap with the heads of our lists
    list := new(ListNode) // Create a new list to hold the merged nodes
    dummy := list // A "dummy" pointer to the head of our new list
    // Keep removing ("Pop" operation in heap) the smallest node from the heap and adding it to the new list
    for len(pq) > 0 {
        n := heap.Pop(&pq)
        node := n.(*ListNode) // node is the list with minimum value in heap
        list.Next = node
        list = list.Next // Update our new list's last node pointer
        // Add the next node of current list to heap (if it's not nil)
        if node.Next != nil {
            heap.Push(&pq, node.Next) 
        }
    }
    // The "Next" attribute of the dummy node gives us the head of the merged list
    return dummy.Next
}

// Here, various methods to manipulate our priority queue (implemented as a heap) are defined.
type PQ []*ListNode

func (pq PQ) Len() int {
    return len(pq)
}

func (pq PQ) Swap(a, b int) {
    pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
    return pq[a].Val < pq[b].Val
}

func (pq *PQ) Pop() (x interface{}) {
    x, *pq = (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
    return x
}

func (pq *PQ) Push(x interface{}) {
    node := x.(*ListNode)
    *pq = append(*pq, node)
}