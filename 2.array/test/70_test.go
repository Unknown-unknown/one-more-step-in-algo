package test

import (
	"fmt"
	"testing"
)

func Test_70(t *testing.T) {
	cases := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 5,
	}
	for k, v := range cases {
		actual := climbStairs(k)
		fmt.Printf("%d stairs expected: %d, actual: %d\n", k, v, actual)
	}
}

func climbStairsBruteForce(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i < n+1; i++ {
		f3 := f1 + f2
		f1 = f2
		f2 = f3
	}
	return f2
}
