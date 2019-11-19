package test

import (
	"fmt"
	"testing"
)

func Test_11(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	max := maxArea(height)
	fmt.Printf("two pointer: %d\n", max)
}

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}
	return max
}
