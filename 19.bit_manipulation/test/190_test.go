package test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_190(t *testing.T) {
	num, _ := strconv.ParseInt("11111111111111111111111111111101", 2, 0)
	fmt.Printf("num: %b\n", num)
	res := reverseBits(uint32(num))
	fmt.Printf("res: %b -> %b\n", num, res)
}

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	t := 32
	for t > 0 {
		ans <<= 1
		ans |= num & 1
		num >>= 1
		t--
		fmt.Printf("for: num: %b, ans: %b, t: %d\n", num, ans, t)
	}
	return ans
}
