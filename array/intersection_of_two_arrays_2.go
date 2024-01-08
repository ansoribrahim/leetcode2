package array

func intersect(nums1 []int, nums2 []int) []int {
	var numMap = make(map[int]int)
	var result []int

	for _, num := range nums1 {
		numMap[num]++
	}

	for _, num := range nums2 {
		if numMap[num] > 0 {
			result = append(result, num)
			numMap[num]--
		}
	}

	return result
}
