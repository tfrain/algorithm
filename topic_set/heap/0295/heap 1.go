// MedianFinder uses two heaps to find the median of the added numbers.
type MedianFinder struct {
    s, b *Heap
}

// Constructor initializes MedianFinder with a min-heap and a max-heap.
func Constructor() MedianFinder {
    return MedianFinder{
        // 's' (min-heap; heap of smaller half) set for ascending order
        s: newHeap(func(i1, i2 int) bool {
            return i1 < i2
        }),
        // 'b' (max-heap; heap of larger half) set for descending order
        b: newHeap(func(i1, i2 int) bool {
            return i1 > i2
        }),
    }
}

// AddNum inserts a number into MedianFinder maintaining the size property of heaps.
func (this *MedianFinder) AddNum(num int) {
    // If total size of both heaps is even, push to max-heap, else push to min-heap
    if (this.s.Len()+this.b.Len())%2 == 0 {
        heap.Push(this.b, num)
        heap.Push(this.s, heap.Pop(this.b))  // Balance the heaps
    } else {
        heap.Push(this.s, num)
        heap.Push(this.b, heap.Pop(this.s))  // Balance the heaps
    }
}

// FindMedian returns the median of all added numbers.
func (this *MedianFinder) FindMedian() float64 {
    // If count of numbers is even, median is the average of top elements from both heaps
    if (this.s.Len()+this.b.Len())%2 == 0 {
        return float64(this.s.Peak()+this.b.Peak()) / 2
    } else {
        // If count of numbers is odd, median is the top element from min-heap
        return float64(this.s.Peak())
    }
}

// Heap is an implementation of a heap data structure with a flexible less function. 
type Heap struct {
    Values   []int  // Values contains the elements of the heap
    LessFunc less  // LessFunc defines the ordering of the heap elements
}

type less func(int, int) bool

func newHeap(l less) *Heap {
	return &Heap{
		LessFunc: l,
	}
}

// Less is a helper function that uses the LessFunc of the heap to compare elements.
func (h Heap) Less(i, j int) bool {
    return h.LessFunc(h.Values[i], h.Values[j])
}

// Swap swaps the elements at indices i and j.
func (h Heap) Swap(i, j int) {
    h.Values[i], h.Values[j] = h.Values[j], h.Values[i]
}

// Len returns the count of elements in the heap.
func (h Heap) Len() int {
    return len(h.Values)
}

// Pop removes and returns the top element from the heap.
func (h *Heap) Pop() (v interface{}) {
    h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
    return v
}

// Peak returns the top element from the heap.
func (h *Heap) Peak() int {
    return h.Values[0]
}

// Push inserts an element at the end of the heap.
func (h *Heap) Push(v interface{}) {
    h.Values = append(h.Values, v.(int))
}