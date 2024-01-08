package array

import "fmt"

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
// solution:
// loop over the array, and skip until found different int
// if found different int, place it to the last different int
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize a pointer to the first element
	// and iterate over the array starting from the second element
	var i int = 1
	for _, num := range nums[1:] {
		// nums[i-1] means start from the 0 index of array,
		// this will compare with the second index (num start from second index)
		// num will be loop until end of array
		if num != nums[i-1] {
			// If the current number is different from the previous one,
			// update the array in-place
			nums[i] = num
			i++
		}
	}

	fmt.Println(nums)
	// Modify the original slice in-place and return the length of the modified array
	return i
}
