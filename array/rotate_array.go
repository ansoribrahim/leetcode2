package array

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
// solution :
// https://www.youtube.com/watch?v=BHr381Guz3Y
// code defided into 3 section :
// section one : reverse array completely
// section two : reversed index 0 to k
// section three : reversed index k+1 to len(nums)

// below when i run nums= 123456, k =2
// and below i separate it to 3 section too

//	123456 l = 0, r = 5
//	623451 l = 1, r = 4
//	653421 l = 2, r = 3
//	654321 l = 3, r = 2
//
//	654321 l = 0, r = 1
//	564321 l = 1, r = 0
//
//	564321 l = 2, r = 5
//	561324 l = 3, r = 4
//	561234 l = 4, r = 3

func rotate(nums []int, k int) {
	// why this be modulo, explanation below
	k = k % len(nums)

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	l, r = 0, k-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	l, r = k, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}
}

//The problem is about rotating an array to the right by k steps.
//If k is greater than the length of the array, it means you might end up rotating the array more than one full circle,
//and you only need to perform the effective rotations.
//
//Let's say the length of the array is n, and you want to rotate it by k steps.
//If k is greater than or equal to n, you can think of it as making several full rotations before reaching the desired
//position. But after each full rotation, you end up in the same state as the initial array.
//
//For example, if n is 5 and k is 7, rotating the array by 7 steps is the same as rotating it by 2 steps (7 % 5 = 2).
//This is because after making two full rotations, you end up in the same position as you would be after
//rotating it by 7 steps.
//
//So, by using k = k % len(nums), you are essentially finding the remainder when k is divided by the length of
//the array. This ensures that k is effectively reduced to the range [0, len(nums)), making sure that you perform only
//the necessary rotations to achieve the desired result without unnecessary full circles.
