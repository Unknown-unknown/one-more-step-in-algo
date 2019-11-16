package test

import (
	"fmt"
	"testing"
)

func Test_120(t *testing.T) {
	triangle := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
	res := minimumTotal(triangle)
	fmt.Printf("1.res: %v\n", res)
}

func TestMin(t *testing.T) {
	arr := []int{1, 4, -3, 0}
	index, min := minimum(arr)
	fmt.Printf("index: %d, min: %d\n", index, min)
}

// ! 错误解法：使用贪心算法，每次取下一层相邻的最小值，但最终求和不一定最小┓( ´∀` )┏
// Testcase: [[-1],[2,3],[1,-1,-3]], Answer: 0, Expected Answer: -1
func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	if len(triangle[0]) <= 0 {
		return 0
	}
	level := len(triangle)
	sum := triangle[0][0]
	tmp := 0
	for i := 1; i < level; i++ {
		row := triangle[i]
		if tmp+1 < len(row) {
			if row[tmp] < row[tmp+1] {
				sum += row[tmp]
			} else {
				sum += row[tmp+1]
				tmp++
			}
		} else {
			sum += row[tmp]
		}
	}
	return sum
}

func minimum(arr []int) (index, min int) {
	if len(arr) == 0 {
		return -1, 0
	}
	min = arr[0]
	index = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}
	return
}
