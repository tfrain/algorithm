func maxSlidingWindow(nums []int, k int) []int {
 n := len(nums)
 res := make([]int, n-k+1) // Array to store maxima
 deQueue := make([]int, 0) // Our DeQueue data structure
 for i, v := range nums {
  // Check if the element at deque front is out of window
  if len(deQueue) != 0 && deQueue[0] < i-k+1 {
   deQueue = deQueue[1:] // Dequeue from front
  }
  // Remove elements from the back of deque which are less than new element v
  for len(deQueue) != 0 && v > nums[deQueue[len(deQueue)-1]] {
   deQueue = deQueue[:len(deQueue)-1] // Dequeue from back
  }
  // Append new element's index to the deque
  deQueue = append(deQueue, i)
  // If the window has reached size k, record the maximum (the value at deque front)
  if i-k+1 >= 0 {
   res[i-k+1] = nums[deQueue[0]]
  }
 }
 return res
}