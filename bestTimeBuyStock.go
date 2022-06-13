package main

func bestTimeBuyStock(prices []int) int {
	leftPointer := 0
	rightPointer := 1
	bestProfit := 0

	for i := 0; i < len(prices)-1; i++ {

		if prices[leftPointer] > prices[rightPointer] {
			leftPointer = rightPointer
		} else {
			// profitable
			profit := prices[rightPointer] - prices[leftPointer]
			if profit > bestProfit {
				bestProfit = profit
			}
		}
		rightPointer += 1

	}
	return bestProfit

}
