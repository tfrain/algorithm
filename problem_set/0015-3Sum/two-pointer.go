package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)                  // Sorts the array, which is essential for using the two-pointer technique
	return nSumTarget(nums, 3, 0, 0) // Calls nSumTarget function to find all triplets that meet the conditions
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	sz := len(nums)
	var ret [][]int
	if n < 2 || sz < start { // Base case checks for invalid input or conditions
		return ret
	}
	if n == 2 { // When n is 2, it's a two-sum problem which can be solved using a two-pointer approach
		lo, hi := start, sz-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++ // Move the left pointer to the right to increase the sum
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi-- // Move the right pointer to the left to decrease the sum
				}
			} else {
				ret = append(ret, []int{left, right}) // Found a valid pair, add to the result
				// Skip duplicates
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else { // For n > 2, recursively reduce the problem to a two-sum problem
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				ret = append(ret, arr) // Append the current number to each result returned by the recursive call
			}
			// Skip duplicates
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ret
}
