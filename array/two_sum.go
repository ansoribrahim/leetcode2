package array

func twoSum(nums []int, target int) []int {
	var lMap = make(map[int]int, len(nums))

	for i, num := range nums {
		resultWant := target - num

		if value, ok := lMap[resultWant]; ok {
			return []int{value, i}
		}
		lMap[num] = i
	}

	return nil
}
