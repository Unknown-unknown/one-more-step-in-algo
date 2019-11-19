package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_66(t *testing.T) {
	original := map[string][]int{
		"no.1:": []int{1, 2, 3},
		"no.2:": []int{1, 2, 9},
		"no.3:": []int{9, 9, 9},
		"no.4:": []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
		"no.5:": []int{0},
	}
	expected := map[string][]int{
		"no.1:": []int{1, 2, 4},
		"no.2:": []int{1, 3, 0},
		"no.3:": []int{1, 0, 0, 0},
		"no.4:": []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		"no.5:": []int{1},
	}
	for k, v := range original {
		res := plusOneBruteForce(v)
		fmt.Printf("%s brute force: %v -> %v, expected: %v\n", k, v, res, expected[k])

		res = plusOneOptimal(v)
		fmt.Printf("%s optimal: %v -> %v, expected: %v\n", k, v, res, expected[k])
	}
}

func plusOneBruteForce(digits []int) []int {
	str := []string{}
	for _, d := range digits {
		str = append(str, strconv.Itoa(d))
	}
	d, err := strconv.Atoi(strings.Join(str, ""))
	if err != nil {
		fmt.Printf("strconv.Atoi err: %v", err)
		return []int{}
	}
	d = d + 1
	plusedStr := strconv.Itoa(d)
	str = strings.Split(plusedStr, "")
	res := []int{}
	for _, s := range str {
		d, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("strconv.Atoi err: %v", err)
			return []int{}
		}
		res = append(res, d)
	}
	return res
}

func plusOneOptimal(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}
