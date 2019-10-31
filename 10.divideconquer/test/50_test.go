package test

import (
	"testing"

	"gotest.tools/assert"
)

func Test_50(t *testing.T) {
	res := myPowDC(5, 0)
	assert.Equal(t, res, 1.0)

	res = myPowDC(2, 3)
	assert.Equal(t, res, 8.0)

	res = myPowDC(2, 4)
	assert.Equal(t, res, 16.0)

	res = myPowNewton(5, 0)
	assert.Equal(t, res, 1.0)

	res = myPowNewton(2, 3)
	assert.Equal(t, res, 8.0)

	res = myPowNewton(2, 4)
	assert.Equal(t, res, 16.0)
}

func myPowDC(x float64, n int) float64 {
	// terminator
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// process current level
	// drill down
	res := myPowDC(x, n/2)
	// merge
	if n%2 != 0 {
		res = res * res * x
	} else {
		res = res * res
	}
	return res
}

func myPowNewton(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	subRes := x
	res := 1.0
	for i := n; i > 0; i = i / 2 {
		if i%2 == 1 {
			res = subRes * res
		}
		subRes = subRes * subRes
	}
	return res
}
