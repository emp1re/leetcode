package test

import (
	"testing"

	"github.com/emp1re/leetcode/other/internal"
)

func TestPrefixSum(t *testing.T) {
	prefix := internal.PrefixSum([]int{1, 2, 3, 4, 5})
	t.Log(prefix)

}
func TestPrefixSumInterval(t *testing.T) {
	prefix := internal.PrefixSum([]int{1, 2, 3, 4, 5})
	t.Log(internal.PrefixSumInterval(prefix, 1, 3))
}
