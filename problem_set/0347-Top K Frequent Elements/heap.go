// topKFrequent function returns the k most frequent elements in nums.
func topKFrequent(nums []int, k int) []int {
    // Initialize the result slice and count map to count frequency of each element.
    var res []int
    cntMap := make(map[int]int)

    // Frequency count of each element
    for _, v := range nums {
        cntMap[v]++
    }

    // Initialize a heap of elements with their count.
    h := &Heap{}

    for key, v := range cntMap {
        e := element{
            val: key,
            cnt: v,
        }

        // Pushing each element and its frequency into the heap.
        heap.Push(h, e)

        // If heap size goes beyond k, popping out the least frequent one.
        if h.Len() > k {
            heap.Pop(h)
        }
    }

    // Popping remaining elements in heap and appending them to the result.
    for h.Len() > 0 {
        res = append(res, heap.Pop(h).(element).val)
    }

    // Return the result slice with k most frequent elements from nums.
    return res
}

// Defining struct to hold element and its count.
type element struct {
    val, cnt int
}

// Defining Heap as slice of elements.
type Heap []element

// Function to return len of heap.
func (h Heap) Len() int {
    return len(h)
}

// Function to swap elements at index i and j.
func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

// Function returning true if frequency of element at i is less than at j.
func (h Heap) Less(i, j int) bool {
    return h[i].cnt < h[j].cnt
}

// Function to push an element into heap.
func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(element))
}

// Function to pop out an element from heap.
func (h *Heap) Pop() (x interface{}) {
    *h, x = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return
}