package array

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 || nums[0] != nums[1] {
		return nums[0]
	}

	if nums[len(nums)-2] != nums[len(nums)-1] {
		return nums[len(nums)-1]
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return 0
}

// watch : https://www.youtube.com/watch?v=qMPX1AOa83k
func singleNumberUsingXor(nums []int) int {
	res := 0
	for _, num := range nums {
		res = num ^ res
	}
	return res
}
