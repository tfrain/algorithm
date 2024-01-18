// findKthLargest uses a min heap to find the k-th largest element in an array.
func findKthLargest(nums []int, k int) int {
    // Create a MinHeap from the array.
    minHeap := MinHeap(nums)
    // Initialize the heap in-place.
    heap.Init(&minHeap)
    // Pop elements from the heap until its size is equal to k.
    for minHeap.Len() > k {
        heap.Pop(&minHeap)
    }
    // The k-th largest element is now at the root of the heap.
    return heap.Pop(&minHeap).(int)
}

// MinHeap is a type that implements the heap.Interface for minimum heap.
type MinHeap []int

// Len returns the number of elements in the heap.
func (h MinHeap) Len() int           { return len(h) }
// Less reports whether the element with index i should sort before the element with index j.
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// Swap swaps the elements with indexes i and j.
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap. The complexity is O(log n) due to the reordering.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element (root) from the heap.
// The complexity is O(log n) as it requires reordering the heap after removal.
func (h *MinHeap) Pop() interface{} {
    n := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return n
}
