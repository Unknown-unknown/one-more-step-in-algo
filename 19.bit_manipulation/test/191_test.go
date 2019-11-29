package test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_191(t *testing.T) {
	num := uint32(00000000000000000000000000001011)
	res := hammingWeight(num)
	fmt.Printf("res: %d\n", res)
}

func hammingWeight(num uint32) int {
	str := strconv.FormatInt(int64(num), 2)
	count := 0
	for _, v := range []rune(str) {
		if v == '1' {
			count++
		}
	}
	return count
}
