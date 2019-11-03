package test

import (
	"fmt"
	"testing"
)

func Test_122(t *testing.T) {
	case1 := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(case1)
	fmt.Printf("1.res: %v\n", res)

	case2 := []int{1, 2, 3, 4, 5}
	res = maxProfit(case2)
	fmt.Printf("2.res: %v\n", res)

	case3 := []int{7, 6, 4, 3, 1}
	res = maxProfit(case3)
	fmt.Printf("3.res: %v\n", res)
}

func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
