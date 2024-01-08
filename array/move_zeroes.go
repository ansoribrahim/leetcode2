package array

func moveZeroes(nums []int) {
	leftPointer := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[leftPointer], nums[i] = nums[i], nums[leftPointer]
			leftPointer++
		}
	}
}
