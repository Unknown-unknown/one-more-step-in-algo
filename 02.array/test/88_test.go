package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_88(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	mergeSortFirst(nums1, m, nums2, n)

	nums1 = []int{1, 2, 3, 0, 0, 0}
	m = 3
	nums2 = []int{2, 5, 6}
	n = 3
	merge(nums1, m, nums2, n)
}

func mergeSortFirst(nums1 []int, m int, nums2 []int, n int) {
	tmp := append(nums1[:m], nums2...)
	sort.Ints(tmp)
	for i := 0; i < m+n; i++ {
		nums1[i] = tmp[i]
	}
	fmt.Printf("res: %v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			fmt.Printf("nums1[%d] = nums1[%d], %v\n", k, i, nums1)
			i--
		} else {
			nums1[k] = nums2[j]
			fmt.Printf("nums1[%d] = nums2[%d], %v\n", k, j, nums1)
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	fmt.Printf("res: %v\n", nums1)
}
