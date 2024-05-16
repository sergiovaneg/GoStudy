package p2997

import (
	"math/bits"
)

func MinOperations(nums []int, k int) int {
	xor := nums[0]
	for _, num := range nums[1:] {
		xor = xor ^ num
	}

	return bits.OnesCount(uint(xor ^ k))
}
