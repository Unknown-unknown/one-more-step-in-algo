package main

import "fmt"

func main() {
	nums := []int{8, 9, 5, 2, 4, 1, 0}
	res := selectionSort(nums)
	fmt.Printf("selectionSort: %v\n", res)

	res = insertionSort(nums)
	fmt.Printf("insertionSort: %v\n", res)

	res = bubbleSort(nums)
	fmt.Printf("bubbleSort: %v\n", res)
}

func selectionSort(nums []int) []int {
	var midIndex int

	for i := 0; i < len(nums); i++ {
		midIndex = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[midIndex] {
				midIndex = j
			}
		}

		nums[midIndex], nums[i] = nums[i], nums[midIndex]
	}

	return nums
}

func insertionSort(nums []int) []int {
	var n = len(nums)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j = j - 1
		}
	}
	return nums
}

func bubbleSort(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
