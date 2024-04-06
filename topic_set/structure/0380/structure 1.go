type RandomizedSet struct {
    nums     []int
    valToIdx map[int]int
}

// Constructs and initializes a new RandomizedSet object
func Constructor() RandomizedSet {
    return RandomizedSet{
        // Initialize nums as an empty slice
        nums: make([]int, 0),
        // Initialize valToIdx as an empty map
        valToIdx: make(map[int]int),
    }
}

// Inserts a value to the set. Returns true if the set did not already contain the specified element
func (this *RandomizedSet) Insert(val int) bool {
    // If the value already exists in the map, return false
    if _, ok := this.valToIdx[val]; ok {
        return false
    }
    // Otherwise, map the value to the current length of nums (which is the index where the value will be appended)
    this.valToIdx[val] = len(this.nums)
    // And append the value to nums
    this.nums = append(this.nums, val)
    
    return true
}

// Removes a value from the set. Returns true if the set contained the specified element
func (this *RandomizedSet) Remove(val int) bool {
    // If the value does not exist in the map, return false
    if _, ok := this.valToIdx[val]; !ok {
        return false
    }
    // Get the index of the value to be removed
    idx := this.valToIdx[val]
    // Update the mapping of the value at the end of nums to this index
    this.valToIdx[this.nums[len(this.nums)-1]] = idx
    // Remove the mapping of the value to be removed
    delete(this.valToIdx, val)
    // Swap the value to be removed with the value at the end of nums
    this.nums[idx], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[idx]
    // Truncate nums to remove the last element
    this.nums = this.nums[:len(this.nums)-1]

    return true
}

// Returns a random element from the current set of elements.
func (this *RandomizedSet) GetRandom() int {
    // Use the rand function to get a random index
    return this.nums[rand.Intn(len(this.nums))]
}