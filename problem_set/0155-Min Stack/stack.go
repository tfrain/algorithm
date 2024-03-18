// MinStack type: an array-based implementation
type MinStack struct {
    min   int     // stores the minimum element present in the stack at any moment
    stack []int   // golang slice working as the stack
}

// Constructor initializes the MinStack
func Constructor() MinStack {
    return MinStack{
        stack: []int{},   // initialize the stack as an empty slice
        min:   math.MaxInt32, // set min to the maximum integer, as defined in the "math" package
    }
}

// Push adds an element to the top of the MinStack
func (this *MinStack) Push(val int) {
    if val < this.min {
        this.min = val   // update min if the pushed value is smaller than the current min
    }
    this.stack = append(this.stack, val)   // add the element to top of the stack
}

// Pop removes the element on top of the stack
func (this *MinStack) Pop() {
    v := this.stack[len(this.stack)-1]   // get top element of stack
    this.stack = this.stack[:len(this.stack)-1]   // remove the top element from the stack
    // if the popped element is min, iterate through the stack to find a new min
    if len(this.stack) != 0 && v == this.min {
        min := this.stack[0]
        for _, val := range this.stack {
            if min > val {
                min = val
            }
        }
        this.min = min
    }
    // if the stack is empty, reset min to MaxInt32
    if len(this.stack) == 0 {
        this.min = math.MaxInt32
    }
}

// Top function returns the top element of the stack without removing it
func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]  // return the top element of the stack
}

// GetMin function returns the minimum element in the MinStack
func (this *MinStack) GetMin() int {
    return this.min  // return the current min
}

// Defining the Node struct for the linked list
type Node struct {
    min  int  // maintains the `min` so far
    val  int  // value for current node
    Next *Node  // pointer to the next node
}

// Defining the MinStack struct using Node as linked list
type MinStack struct {
    n *Node  // head of the linked list
}

// Constructor initializes an empty MinStack
func Constructor() MinStack {
    return MinStack{}
}

// Push adds an element to the top of the stack
func (this *MinStack) Push(val int) {
    min := val  // initialize `min` as the value to be pushed
    head := this.n  // get the head of the linked list
    if head != nil && min > head.min {
        min = head.min  // update `min` if `head.min` is smaller
    }
    node := &Node{
        val:  val,  // value for the new node
        min:  min,  // `min` for the new node is the min of `head.min` and val
        Next: this.n,  // new node's next is the old head, new node becomes the new head
    }
    this.n = node  // update the head of linked list to the new node
}

// Pop removes the top element from the stack
func (this *MinStack) Pop() {
    this.n = this.n.Next  // move the head to the next node in the linked list
}

// Top returns the top element from the stack
func (this *MinStack) Top() int {
    return this.n.val  // return the value of the head of linked list
}

// GetMin returns the minimum element in the stack
func (this *MinStack) GetMin() int {
    return this.n.min  // return the `min` value of the head of linked list
}