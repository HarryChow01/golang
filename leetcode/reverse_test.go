package leetcode

import (
	"math"
	"testing"
)

func reverseInt(x int) int {
	var n int64 = 0
	for x != 0 {
		n = n*int64(10) + int64(x%10)
		x /= 10
	}

	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}
	return int(n)
}

func TestReverseInt(t *testing.T) {
	n := reverseInt(654)
	t.Log(n)
}
