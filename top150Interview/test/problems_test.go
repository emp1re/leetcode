package test

import (
	"testing"

	"github.com/emp1re/leetcode/top150/internal/problems"
)

func TestProblems(t *testing.T) {

}
func BenchmarkMinimumSizeSubarraySum(b *testing.B) {
	bigSet := []int{2, 3, 1, 2, 4, 3}
	b.SetBytes(2)
	test := problems.MinSubArrayLen(396893380, bigSet)

	b.Log("result :", test)
}
