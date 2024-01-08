package array

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
func maxProfit(prices []int) int {
	profit := 0

	//maxPrice := 10000 //max from requirement
	//minPrice := maxPrice
	//for i, price := range prices {
	//	if price < minPrice {
	//		minPrice = price
	//	} else {
	//		notLastIndex := len(prices)-1 > i
	//
	//		if notLastIndex && prices[i] < prices[i+1] {
	//			continue
	//		} else {
	//			profit = profit + (price - minPrice)
	//			minPrice = maxPrice
	//		}
	//	}
	//}

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
